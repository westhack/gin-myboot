package api

import (
	"gin-myboot/global"
	"gin-myboot/middleware"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/common/service"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	systemRes "gin-myboot/modules/system/model/response"
	"gin-myboot/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type AuthApi struct {
}

var authService = service.CommonServiceGroup.AuthService

// Login
// @Tags Auth
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/common/auth/login [post]
func (b *AuthApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l)

	if l.Mobile != "" {
		b.mobile(c, l)
	} else {
		b.password(c, l)
	}
}

func (b *AuthApi) mobile(c *gin.Context, l systemReq.Login) {
	if err, _ := utils.CheckSmsCode(l.Mobile, l.SmsCode, true); err == nil {
		if err, user := authService.Mobile(l.Mobile); err != nil {
			global.Error("登陆失败! 手机号未注册!", err)
			response.FailWithMessage("手机号未注册", c)
		} else {
			b.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("手机验证码错误"+err.Error(), c)
	}
}

func (b *AuthApi) password(c *gin.Context, l systemReq.Login) {
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := authService.Login(u); err != nil {
			global.Error("登陆失败! 用户名不存在或者密码错误!", err)
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			b.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (b *AuthApi) tokenNext(c *gin.Context, user system.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	var roleId uint64
	var roleName string

	roleName = user.Role.Name
	roleId = user.Role.ID

	claims := systemReq.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Nickname:   user.Nickname,
		Username:   user.Username,
		RoleId:     roleId,
		RoleName:   roleName,
		BufferTime: global.Config.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Config.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "myBoot",                                          // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Error("获取token失败!", err)
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.Config.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.Error("设置登录状态失败!", err)
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.Error("设置登录状态失败!", err)
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// Register
// @Tags Auth
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /v1/common/auth/register [post]
func (b *AuthApi) Register(c *gin.Context) {
	var r systemReq.Register
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("注册失败，"+utils.GetError(err, r), c)
		return
	}

	var roles []system.SysRole
	var roleId uint64
	_ = global.GormDB.Where("default_role = 1").Find(&roles).Error
	if len(roles) > 0 {
		roleId = roles[0].ID
	}

	if r.Nickname == "" {
		r.Nickname = r.Username
	}

	user := &system.SysUser{
		Username: r.Username, Nickname: r.Nickname, Mobile: r.Mobile,
		Password: r.Password, Avatar: r.Avatar,
		RoleId: roleId, Roles: roles,
		Status: 2, IsAdmin: false, Type: int(0),
	}

	err, _ := authService.Register(*user)
	if err != nil {
		global.Error("注册失败!", err)
		response.FailWithMessage("注册失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

// ChangePassword
// @Tags Auth
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.UserChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /v1/common/auth/changePassword [post]
func (b *AuthApi) ChangePassword(c *gin.Context) {
	var user systemReq.UserChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, reqUser := utils.GetCurrUser(c); err != nil {
		return
	} else {
		if reqUser.Password == utils.MD5V([]byte(user.Password)) {
			if err, _ := authService.ChangePassword(&reqUser, user.NewPassword); err != nil {
				global.Error("修改失败!", err)
				response.FailWithMessage("修改失败", c)
			} else {
				response.OkWithMessage("修改成功", c)
			}
		} else {
			response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		}
	}
}

// Update
// @Tags Auth
// @Summary 用户修改资料
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.UserUpdateAccountRequest true "用户名, 邮箱, 昵称，头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /v1/common/auth/update [post]
func (b *AuthApi) Update(c *gin.Context) {
	var reqUser systemReq.UserUpdateAccountRequest
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		response.FailWithMessage("修改失败，"+utils.GetError(err, reqUser), c)
		return
	}

	if err, user := utils.GetCurrUser(c); err != nil {
		return
	} else {
		reqUser.Id = user.ID
		if err, _ := authService.Update(reqUser); err != nil {
			global.Error("修改失败!", err)
			response.FailWithMessage("修改失败", c)
		} else {
			response.OkWithMessage("修改成功", c)
		}
	}
}

// GetUserInfo
// @Tags Auth
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/common/auth/getUserInfo [post]
func (b *AuthApi) GetUserInfo(c *gin.Context) {
	if err, reqUser := utils.GetCurrUser(c); err != nil {
	} else {
		_, permissions := permissionMenuService.GetUserRolePermissionList(reqUser.RoleId)

		reqUser.Permissions = permissions

		response.OkWithDetailed(gin.H{"userInfo": reqUser}, "获取成功", c)
	}
}

// GetUserMenus
// @Tags Auth
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/v1/common/auth/getUserMenus [post]
func (a *AuthApi) GetUserMenus(c *gin.Context) {
	if err, menus := permissionMenuService.GetUserRolePermissionMenuTree(utils.GetUserRoleId(c)); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysRolePermissionMenu{}
		}
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// Logout
// @Tags Auth
// @Summary 退出登录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/common/auth/logout [post]
func (a *AuthApi) Logout(c *gin.Context) {

	_, user := utils.GetUser(c)

	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		jwtService.JsonInBlacklist(blackJWT)
	}

	response.OkWithDetailed(gin.H{"reload": true}, "退出成功", c)
}
