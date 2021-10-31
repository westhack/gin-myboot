package service

import (
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type RoleService struct {
}

var RoleServiceApp = new(RoleService)

// Create
// @function: Create
// @description:  创建一个角色
// @param: role request.RoleFormRequest
// @return: err error
func (roleService *RoleService) Create(role systemReq.RoleFormRequest) (err error) {

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {

		if !errors.Is(tx.Where(" name = ?", role.Name).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
			global.Debug("CreateRole", "存在相同name修改失败")
			return errors.New("存在相同name修改失败")
		}

		newRole := system.SysRole{}

		newRole.ParentId    = role.ParentId
		newRole.Title       = role.Title
		newRole.Name        = role.Name
		newRole.Description = role.Description
		newRole.DefaultRole = role.DefaultRole

		txErr := tx.Create(&newRole).Error
		if txErr != nil {
			global.Debug("CreateRole", txErr.Error())
			return txErr
		}

		role.Id = newRole.ID
		if len(role.Permissions) > 0 {
			err := roleService.SetRolePermission(tx, role)
			if err != nil {
				return err
			}
		}

		if len(role.Departments) > 0 {
			err := roleService.SetRoleDepartment(tx, role)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// Update
// @function: Update
// @description: 更改一个角色
// @param: role request.RoleFormRequest
// @return: err error
func (roleService *RoleService) Update(role systemReq.RoleFormRequest) (err error) {
	if role.Id == role.ParentId {
		return errors.New("上级节点不能为自己")
	}

	var oldRole system.SysRole

	dataMap := make(map[string]interface{})

	dataMap["parent_id"]   = role.ParentId
	dataMap["title"]       = role.Title
	dataMap["name"]        = role.Name
	dataMap["description"] = role.Description
	dataMap["default_role"] = role.DefaultRole

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", role.Id).First(&oldRole)
		if oldRole.Name != role.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", role.Id, role.Name).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
				global.Debug("UpdateRole", "存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}

		txErr := tx.Unscoped().Delete(&system.SysRolePermission{}, "role_id = ?", role.Id).Error
		if txErr != nil {
			global.Debug("UpdateRole", txErr.Error())
			return txErr
		}

		txErr = tx.Unscoped().Delete(&system.SysRoleDepartment{}, "role_id = ?", role.Id).Error
		if txErr != nil {
			global.Debug("UpdateRole", txErr.Error())
			return txErr
		}

		if len(role.Permissions) > 0 {
			err := roleService.SetRolePermission(tx, role)
			if err != nil {
				return err
			}
		}

		if len(role.Departments) > 0 {
			err := roleService.SetRoleDepartment(tx, role)
			if err != nil {
				return err
			}
		}

		txErr = db.Updates(dataMap).Error
		if txErr != nil {
			global.Debug("UpdateRole", txErr.Error())
			return txErr
		}
		return nil
	})

	return err
}

func (roleService *RoleService) SetRolePermission(tx *gorm.DB, role systemReq.RoleFormRequest) (err error) {
	var permissions []system.SysRolePermission
	if len(role.Permissions) > 0 {
		for _, permission := range role.Permissions {
			if permission > 0 {
				p := system.SysRolePermission{
					RoleId:       role.Id,
					PermissionId: permission,
				}
				permissions = append(permissions, p)
			}
		}
		err = tx.Create(&permissions).Error
		if err != nil {
			global.Debug("SetRolePermission", err.Error())
			return err
		}

		err = roleService.SetRolePermissionApis(tx, role.Id, role.Permissions)
		if err != nil {
			global.Debug("UpdateRole", err.Error())
			return err
		}
	}

	return err
}

func (roleService *RoleService) SetRoleDepartment(tx *gorm.DB, role systemReq.RoleFormRequest) (err error) {
	var departments []system.SysRoleDepartment
	if len(role.Departments) > 0 {
		for _, department := range role.Departments {
			if department > 0 {
				p := system.SysRoleDepartment{
					RoleId:       role.Id,
					DepartmentId: department,
				}
				departments = append(departments, p)
			}
		}

		txErr := tx.Create(&departments).Error
		if txErr != nil {
			global.Debug("SetRolePermission", txErr.Error())
			return txErr
		}
	}

	return err
}

// SetRolePermissionApis
// @function: SetRolePermissionApis
// @param: GormDB *gorm.DB, roleId uint, ids []int
// @description: 设置角色权限
// @return: err error
func (roleService *RoleService) SetRolePermissionApis(GormDB *gorm.DB, roleId uint64, ids []uint64) (err error) {

	if len(ids) == 0 {
		return err
	}

	if roleId == 0 {
		return err
	}

	var allPermissions []system.SysPermission
	err = GormDB.Where("id in (?)", ids).Find(&allPermissions).Error
	if err != nil {
		return err
	}

	var casbinInfos []systemReq.CasbinInfo
	for _, permission := range allPermissions {
		if permission.Api == "" {
			continue
		}

		apis := strings.Split(permission.Api, ",")
		for _, api := range apis {
			strArr := strings.Split(api, ":")
			var path string
			var method string
			if len(strArr) >= 2 {
				path = strArr[0]
				method = strArr[1]
			} else {
				path = strArr[0]
				method = ""
			}

			isExist := false
			for _, info := range casbinInfos {
				if info.Path == path && info.Method == method {
					isExist = true
				}
			}

			if isExist != true {
				casbin := systemReq.CasbinInfo{
					Path:   path,
					Method: method,
				}
				casbinInfos = append(casbinInfos, casbin)
			}
		}
	}

	if len(casbinInfos) == 0 {
		return err
	}
	err = CasbinServiceApp.UpdateCasbin(roleId, casbinInfos)

	return err
}

// Delete
// @function: Delete
// @description: 删除角色
// @param: req *request.GetById
// @return: err error
func (roleService *RoleService) Delete(id uint64) (err error) {
	var role system.SysRole
	err = global.GormDB.Preload("Permissions").Preload("Departments").Where("id = ?", id).First(&role).Error
	if err != nil {
		return err
	}

	if !errors.Is(global.GormDB.Where("role_id = ?", id).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New(role.Name + " 此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GormDB.Where("role_id = ?", id).First(&system.SysUserRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New(role.Name + "此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GormDB.Where("parent_id = ?", id).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New(role.Name + "此角色存在子角色不允许删除")
	}

	err = global.GormDB.Unscoped().Where("id = ?", id).Delete(&system.SysRole{}).Error

	if err != nil {
		return err
	}

	if len(role.Permissions) > 0 {
		err = global.GormDB.Model(role).Association("Permissions").Delete(role.Permissions)
	}

	if len(role.Departments) > 0 {
		err = global.GormDB.Model(role).Association("Departments").Delete(role.Departments)
	}

	err = global.GormDB.Delete(&[]system.SysUserRole{}, "role_id = ?", role.ID).Error

	roleIdStr := strconv.FormatUint(role.ID, 10)
	CasbinServiceApp.ClearCasbin(0, roleIdStr)

	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 删除角色
// @param: req *request.GetByIds
// @return: err error
func (roleService *RoleService) DeleteByIds(ids []uint64) (err error) {
	for _, id := range ids {
		err = roleService.Delete(id)
		if err != nil {
			return err
		}
	}

	return err
}

// GetList
// @function: GetList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (roleService *RoleService) GetList() (err error, list interface{}, total int64) {

	db := global.GormDB
	var role []system.SysRole
	err = db.Preload("Permissions").Preload("RolePermissions").
		Preload("Departments").Preload("RoleDepartments").
		Where("parent_id = 0").Find(&role).Error
	if len(role) > 0 {
		for k := range role {
			err = roleService.findChildrenRole(&role[k])
		}
	}
	return err, role, total
}

// GetRoleInfo
// @function: GetRoleInfo
// @description: 获取所有角色信息
// @param: auth model.SysRole
// @return: err error, sa model.SysRole
func (roleService *RoleService) GetRoleInfo(role system.SysRole) (err error, sa system.SysRole) {
	err = global.GormDB.Preload("Permissions").Preload("RolePermissions").
		Preload("Departments").Preload("RoleDepartments").
		Where("id = ?", role.ID).First(&sa).Error
	return err, sa
}

// SetDataRole
// @function: SetDataRole
// @description: 设置角色资源权限
// @param: auth model.SysRole
// @return: error
func (roleService *RoleService) SetDataRole(role system.SysRole) error {
	var s system.SysRole
	global.GormDB.Preload("RoleDepartments").First(&s, "id = ?", role.ID)
	err := global.GormDB.Model(&s).Association("RoleDepartments").Replace(&role.RoleDepartments)
	return err
}

// SetPermissionMenuRole
// @function: SetPermissionMenuRole
// @description: 菜单权限与角色绑定
// @param: auth *model.SysRole
// @return: error
func (roleService *RoleService) SetPermissionMenuRole(role *system.SysRole) error {
	var s system.SysRole
	global.GormDB.Preload("Permissions").First(&s, "id = ?", role.ID)
	err := global.GormDB.Model(&s).Association("Permissions").Replace(&role.Permissions)
	return err
}

//@function: findChildrenRole
//@description: 查询子角色
//@param: role *model.SysRole
//@return: err error
func (roleService *RoleService) findChildrenRole(role *system.SysRole) (err error) {
	err = global.GormDB.
		Preload("Permissions").Preload("RolePermissions").
		Preload("Departments").Preload("RoleDepartments").
		Where("parent_id = ?", role.ID).Find(&role.Children).Error
	if len(role.Children) > 0 {
		for k := range role.Children {
			err = roleService.findChildrenRole(&role.Children[k])
		}
	}
	return err
}

// GetAll
// @function: GetAll
// @description: 获取全部角色
// @param: role *model.SysRole
// @return: err error, list interface{}, total int64
func (roleService *RoleService) GetAll() (err error, list interface{}, total int64) {
	db := global.GormDB
	var roles []system.SysRole
	err = db.Find(&roles).Error

	return err, roles, int64(len(roles))
}
