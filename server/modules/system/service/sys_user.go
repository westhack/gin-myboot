package service

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

// ChangePassword
// @function: ChangePassword
// @description: 修改用户密码
// @param: u *model.SysUser, newPassword string
// @return: err error, userInter *model.SysUser
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GormDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// Delete
// @function: Delete
// @description: 删除用户
// @param: id uint64
// @return: err error
func (userService *UserService) Delete(id uint64) (err error) {
	var user system.SysUser
	err = global.GormDB.Where("id = ?", id).Delete(&user).Error
	err = global.GormDB.Delete(&[]system.SysUserRole{}, "user_id = ?", id).Error
	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 批量删除用户
// @param: id []uint64
// @return: err error
func (userService *UserService) DeleteByIds(id []uint64) (err error) {
	var user system.SysUser
	err = global.GormDB.Where("id in ?", id).Delete(&user).Error
	err = global.GormDB.Delete(&[]system.SysUserRole{}, "user_id = ?", id).Error
	return err
}

// GetUserInfo
// @function: GetUserInfo
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: err error, user system.SysUser
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.GormDB.Preload("Roles").Preload("Role").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

// FindById
// @function: FindById
// @description: 通过id获取用户信息
// @param: id uint64
// @return: err error, user *model.SysUser
func (userService *UserService) FindById(id uint64) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.GormDB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// FindByUuid
// @function: FindByUuid
// @description: 通过uuid获取用户信息
// @param: uuid string
// @return: err error, user *model.SysUser
func (userService *UserService) FindByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.GormDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

// GetList
// @function: GetList
// @description: 分页获取数据
// @param: info request.QueryParams
// @return: err error, list interface{}, total int64
func (userService *UserService) GetList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)
	db := global.GormDB.Model(&system.SysUser{})
	var userList []system.SysUser

	if queryParams.SortOrder.Column == "" {
		queryParams.SortOrder.Column = "id"
		queryParams.SortOrder.Order = "desc"
	}

	db.Where("is_admin = 0").Scopes(model.Search(queryParams.Search))

	err = db.Count(&total).Error

	err = db.Scopes(model.SortOrder(queryParams.SortOrder)).Limit(limit).Offset(offset).
		Preload("Roles").Preload("Role").Find(&userList).Error

	return err, userList, total
}

// Create
// @function: Create
// @description: 创建用户
// @param: reqUser systemReq.UserFormRequest
// @return: err error, userInter system.SysUser
func (userService *UserService) Create(reqUser *systemReq.UserFormRequest) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.GormDB.Where("username =?", reqUser.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已存在"), userInter
	}
	if !errors.Is(global.GormDB.Where("mobile =?", reqUser.Mobile).First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册
		return errors.New("手机号已存在"), userInter
	}
	if !errors.Is(global.GormDB.Where("email =?", reqUser.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已存在"), userInter
	}

	user.Email = reqUser.Email
	user.Username = reqUser.Username
	user.Mobile = reqUser.Mobile
	user.Nickname = reqUser.Nickname
	user.Password = utils.MD5V([]byte(reqUser.Password))
	user.Status = reqUser.Status
	user.IsAdmin = false
	user.UUID = uuid.NewV4()

	err = global.GormDB.Create(&user).Error
	return err, user
}

// Update
// @function: Update
// @description: 修改用户
// @param: reqUser systemReq.UpdateUserRequest
// @return: err error, user system.SysUser
func (userService *UserService) Update(reqUser *systemReq.UserFormRequest) (err error, user system.SysUser) {
	err = global.GormDB.First(&user, "id = ?", reqUser.Id).Error
	if err != nil {
		return errors.New("用户不存在"), user
	}

	if !errors.Is(global.GormDB.Where("username =?", reqUser.Username).Where("id !=?", reqUser.Id).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已存在"), user
	}
	if !errors.Is(global.GormDB.Where("mobile =?", reqUser.Mobile).Where("id !=?", reqUser.Id).First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册
		return errors.New("手机号已存在"), user
	}
	if !errors.Is(global.GormDB.Where("email =?", reqUser.Email).Where("id !=?", reqUser.Id).First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册
		return errors.New("邮箱已存在"), user
	}

	user.Email = reqUser.Email
	user.Username = reqUser.Username
	user.Mobile = reqUser.Mobile
	user.Nickname = reqUser.Nickname
	user.Status = reqUser.Status

	if reqUser.Password != "" {
		user.Password = utils.MD5V([]byte(reqUser.Password))
	}

	err = global.GormDB.Model(&user).Where("id = ?", reqUser.Id).Updates(user).Error
	return err, user
}

// ResetPassword
// @function: ResetPassword
// @description: 重置密码为 123456
// @param: id []uint64
// @return: err error
func (userService *UserService) ResetPassword(id []uint64) (err error) {
	password := utils.MD5V([]byte("123456"))
	err = global.GormDB.Model(&system.SysUser{}).Where("id in ?", id).Update("password", password).Error
	return err
}

// ChangeStatus
// @function: ChangeStatus
// @description: 修改用户状态
// @param: id []uint64, status int
// @return: err error
func (userService *UserService) ChangeStatus(id []uint64, status int) (err error) {
	err = global.GormDB.Model(&system.SysUser{}).Where("id in ?", id).Update("status", status).Error
	return err
}
