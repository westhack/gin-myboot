package service

import (
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AuthService struct {
}

// Register
// @function: Register
// @description: 用户注册
// @param: u model.SysUser
// @return: err error, userInter model.SysUser
func (authService *AuthService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.GormDB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	if !errors.Is(global.GormDB.Where("mobile = ?", u.Mobile).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("手机号已注册"), userInter
	}
	if !errors.Is(global.GormDB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("邮箱已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GormDB.Create(&u).Error
	return err, u
}

// Login
// @function: Login
// @description: 用户登录
// @param: u *model.SysUser
// @return: err error, userInter *model.SysUser
func (authService *AuthService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	global.Error("=====>", u.Password)
	err = global.GormDB.Where("username = ? AND password = ?", u.Username, u.Password).
		Preload("Roles").Preload("Role").First(&user).Error
	return err, &user
}


// Mobile
// @function: Mobile
// @description: 用户登录
// @param: u *model.SysUser
// @return: err error, userInter *model.SysUser
func (authService *AuthService) Mobile(m string) (err error, userInter *system.SysUser) {
	var user system.SysUser

	err = global.GormDB.Where("mobile = ?", m).Preload("Roles").Preload("Role").First(&user).Error

	return err, &user
}

// ChangePassword
// @function: ChangePassword
// @description: 修改用户密码
// @param: u *model.SysUser, newPassword string
// @return: err error, userInter *model.SysUser
func (authService *AuthService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	err = global.GormDB.Where("username = ? ", u.Username).First(&user).
		Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// GetUserInfo
// @function: GetUserInfo
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: err error, user system.SysUser
func (authService *AuthService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.GormDB.Preload("Roles").Preload("Role").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

// Update
// @function: Update
// @description: 修改用户
// @param: reqUser systemReq.UpdateAccountRequest
// @return: err error, user system.SysUser
func (authService *AuthService) Update(reqUser systemReq.UserUpdateAccountRequest) (err error, user system.SysUser) {

	err = authService.CheckUserInfo(reqUser)
	if err != nil  {
		return err, user
	}

	data := make(map[string]interface{})

	if reqUser.Nickname != "" {
		data["nickname"] = reqUser.Nickname
	}
	if reqUser.Avatar != "" {
		data["avatar"] = reqUser.Avatar
	}
	if reqUser.Mobile != "" {
		data["mobile"] = reqUser.Mobile
	}
	if reqUser.Mobile != "" {
		data["mobile"] = reqUser.Mobile
	}
	if reqUser.Email != "" {
		data["email"] = reqUser.Email
	}

	err = global.GormDB.Model(&user).Where("id = ?", reqUser.Id).Updates(data).Error

	return err, user
}

func (authService *AuthService) CheckUserInfo(reqUser systemReq.UserUpdateAccountRequest) (err error) {
	var user system.SysUser

	err = global.GormDB.Where("id != ? and username = ?", reqUser.Id, reqUser.Username).First(&user).Error
	if user.ID > 0 {
		return errors.New("用户名已存在")
	}

	err = global.GormDB.Where("id != ? and mobile = ?", reqUser.Id, reqUser.Mobile).First(&user).Error
	if user.ID > 0 {
		return errors.New("手机号")
	}

	err = global.GormDB.Where("id != ? and email = ?", reqUser.Id, reqUser.Email).First(&user).Error
	if user.ID > 0 {
		return errors.New("邮箱已存在")
	}

	return err
}