package service

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	systemReq "gin-myboot/modules/system/model/request"
	system "gin-myboot/modules/system/model"
	"gin-myboot/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AdminService struct {
}

// GetList
// @function: GetList
// @description: 分页获取数据
// @param: info request.QueryParams
// @return: err error, list interface{}, total int64
func (adminService *AdminService) GetList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)

	db := global.GormDB.Model(&system.SysUser{}).Where("is_admin=1").Scopes(model.Search(queryParams.Search))

	var userList []system.SysUser

	if queryParams.SortOrder.Column == "" {
		queryParams.SortOrder.Column = "id"
		queryParams.SortOrder.Order = "desc"
	}

	err = db.Count(&total).Error

	err = db.Scopes(model.SortOrder(queryParams.SortOrder)).Limit(limit).Offset(offset).
		Preload("Roles").Preload("Role").Preload("Department").Find(&userList).Error

	return err, userList, total
}

// GetAll
// @function: GetList
// @description: 获取全部管理员
// @return: err error, list interface{}, total int64
func (adminService *AdminService) GetAll(limit int) (err error, list interface{}, total int64) {
	var users []system.SysUser
	err = global.GormDB.Where("status = 1 and is_admin = 1").Preload("Roles").Preload("Role").Limit(limit).Find(&users).Error

	return err, users, total
}

// Create
// @function: Create
// @description: 创建管理员
// @param: u request.AdminFormRequest
// @return: err error, userInter model.SysUser
func (adminService *AdminService) Create(reqUser *systemReq.AdminFormRequest) (err error, userInter system.SysUser) {
	var user system.SysUser
	var roles []system.SysRole

	if !errors.Is(global.GormDB.Where("username = ?", reqUser.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已存在"), userInter
	}
	if !errors.Is(global.GormDB.Where("mobile = ?", reqUser.Mobile).First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册
		return errors.New("手机号已存在"), userInter
	}
	if !errors.Is(global.GormDB.Where("email = ?", reqUser.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断邮箱是否已经注册
		return errors.New("邮箱已存在"), userInter
	}

	if reqUser.Password == "" {
		return errors.New("密码不能为空"), userInter
	}

	err = global.GormDB.Where("id in ?", reqUser.Roles).Find(&roles).Error
	if err == nil {
		user.Roles = roles
	}

	if len(user.Roles) > 0 {
		user.RoleId = roles[0].ID
	}

	user.Email = reqUser.Email
	user.Username = reqUser.Username
	user.Mobile = reqUser.Mobile
	user.Nickname = reqUser.Nickname
	user.Password = utils.MD5V([]byte(reqUser.Password))
	user.Status = reqUser.Status
	user.IsAdmin = true
	user.UUID = uuid.NewV4()
	user.DepartmentId = reqUser.DepartmentId

	err = global.GormDB.Create(&user).Error

	return err, user
}

// Update
// @function: Update
// @description: 修改管理员
// @param: reqUser request.AdminFormRequest
// @return: err error, user system.SysUser
func (adminService *AdminService) Update(reqUser *systemReq.AdminFormRequest) (err error, user system.SysUser) {
	err = global.GormDB.First(&user, "id = ?", reqUser.Id).Error
	if err != nil {
		return errors.New("用户不存在"), user
	}
	if !errors.Is(global.GormDB.Where("username =?", reqUser.Username).Where("id !=?", reqUser.Id).
		First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册

		return errors.New("用户名已存在"), user
	}
	if !errors.Is(global.GormDB.Where("mobile =?", reqUser.Mobile).Where("id !=?", reqUser.Id).
		First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册

		return errors.New("手机号已存在"), user
	}
	if !errors.Is(global.GormDB.Where("email =?", reqUser.Email).Where("id !=?", reqUser.Id).
		First(&user).Error, gorm.ErrRecordNotFound) { // 判断邮箱是否注册

		return errors.New("邮箱已存在"), user
	}

	var roles []system.SysRole

	err = global.GormDB.Where("id in ?", reqUser.Roles).Find(&roles).Error

	if err == nil {
		user.Roles = roles
		global.GormDB.Model(&user).Association("Roles").Replace(user.Roles)
	}

	if len(user.Roles) > 0 {
		user.RoleId = roles[0].ID
	}

	user.Email = reqUser.Email
	user.Username = reqUser.Username
	user.Mobile = reqUser.Mobile
	user.Nickname = reqUser.Nickname
	user.Status = reqUser.Status
	user.DepartmentId = reqUser.DepartmentId

	if reqUser.Password != "" {
		user.Password = utils.MD5V([]byte(reqUser.Password))
	}

	err = global.GormDB.Model(&user).Updates(user).Error

	return err, user
}

// Delete
// @function: Delete
// @description: 删除管理员
// @param: id uint64
// @return: err error
func (adminService *AdminService) Delete(id uint64) (err error) {
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
func (adminService *AdminService) DeleteByIds(id []uint64) (err error) {
	var user system.SysUser

	err = global.GormDB.Where("id in ?", id).Delete(&user).Error
	err = global.GormDB.Delete(&[]system.SysUserRole{}, "user_id = ?", id).Error

	return err
}

// ChangePassword
// @function: ChangePassword
// @description: 修改管理员密码
// @param: u *model.SysUser, newPassword string
// @return: err error, userInter *model.SysUser
func (adminService *AdminService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser

	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GormDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).
		Update("password", utils.MD5V([]byte(newPassword))).Error

	return err, u
}

// SetUserRole
// @function: SetUserRole
// @description: 设置一个管理员的权限
// @param: id uint64, uuid uuid.UUID, roleId uint64
// @return: err error
func (adminService *AdminService) SetUserRole(id uint64, uuid uuid.UUID, roleId uint64) (err error) {
	assignErr := global.GormDB.Where("user_id = ? AND role_id = ?", id, roleId).First(&system.SysUserRole{}).Error

	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}

	err = global.GormDB.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("role_id", roleId).Error

	return err
}

// SetUserRoles
// @function: SetUserRoles
// @description: 设置一个管理员的权限
// @param: id uint64, roleIds []uint64
// @return: err error
func (adminService *AdminService) SetUserRoles(id uint64, roleIds []uint64) (err error) {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserRole{}, "user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}

		var useRole []system.SysUserRole
		for _, v := range roleIds {
			useRole = append(useRole, system.SysUserRole{
				UserId: id, RoleId: v,
			})
		}

		TxErr = tx.Create(&useRole).Error
		if TxErr != nil {
			return TxErr
		}

		// 返回 nil 提交事务
		return nil
	})
}

// Find
// @function: Find
// @description: 获取用户信息
// @param: int64
// @return: err error, user system.SysUser
func (adminService *AdminService) Find(id int64) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.GormDB.Preload("Roles").Preload("Role").First(&reqUser, "id = ?", id).Error
	return err, reqUser
}

// FindById
// @function: FindById
// @description: 通过id获取用户信息
// @param: id int64
// @return: err error, user *model.SysUser
func (adminService *AdminService) FindById(id int64) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.GormDB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// FindByUuid
// @function: FindByUuid
// @description: 通过uuid获取用户信息
// @param: uuid string
// @return: err error, user *model.SysUser
func (adminService *AdminService) FindByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.GormDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

// ResetPassword
// @function: ResetPassword
// @description: 重置密码为 123456
// @param: id []uint64
// @return: err error
func (adminService *AdminService) ResetPassword(id []uint64) (err error) {
	password := utils.MD5V([]byte("123456"))
	err = global.GormDB.Model(&system.SysUser{}).Where("id in ?", id).Update("password", password).Error
	return err
}

// ChangeStatus
// @function: ChangeStatus
// @description: 修改用户状态
// @param: id []uint64, status int
// @return: err error
func (adminService *AdminService) ChangeStatus(id []uint64, status int) (err error) {
	err = global.GormDB.Model(&system.SysUser{}).Where("id in ?", id).Update("status", status).Error
	return err
}
