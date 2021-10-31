package service

import (
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"gorm.io/gorm"
	"strings"
)

type PermissionService struct {
}

var PermissionAppService = new(PermissionService)

// Create
// @function: Create
// @description: 添加权限菜单
// @param: permission request.PermissionFormRequest
// @return: error
func (permissionService *PermissionService) Create(permission *request.PermissionFormRequest) (err error) {

	newPermission := system.SysPermission{}

	newPermission.KeepAlive = permission.KeepAlive
	newPermission.ParentId = permission.ParentId
	newPermission.Path = permission.Path
	newPermission.Name = permission.Name
	newPermission.Hidden = permission.Hidden
	newPermission.Component = permission.Component
	newPermission.Title = permission.Title
	newPermission.SortOrder = permission.SortOrder
	newPermission.Redirect = permission.Redirect
	newPermission.IsButton = permission.IsButton
	newPermission.Status = permission.Status
	newPermission.Redirect = permission.Redirect
	newPermission.Api = permission.Api

	if permission.Icon != "" {
		newPermission.Icon = permission.Icon
	} else {
		newPermission.Icon = "menu"
	}
	if permission.Path == "" && permission.Component != "" {
		newPermission.Path = permission.Name
	}

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {

		txErr := tx.Create(&newPermission).Error

		var buttons []system.SysPermission
		if len(permission.Buttons) > 0 {
			for _, button := range permission.Buttons {
				if button != "" {
					s := strings.Split(button, ":")
					if len(s) >= 2 {
						p := system.SysPermission{}

						p.ParentId = newPermission.ID
						p.Name = s[0]
						p.Title = s[1]
						p.Icon = "plus-circle"

						buttons = append(buttons, p)
					}
				}

			}
			txErr := tx.Create(&buttons).Error
			if txErr != nil {
				global.Debug("CreatePermission", txErr.Error())
				return txErr
			}
		}

		if txErr != nil {
			global.Debug("CreatePermission", txErr.Error())
			return txErr
		}
		return nil
	})

	return err
}

// Delete
// @function: Delete
// @description: 删除权限菜单
// @param: id float64
// @return: err error
func (permissionService *PermissionService) Delete(id uint64) (err error) {
	err = global.GormDB.Where("parent_id = ?", id).First(&system.SysPermission{}).Error
	if err != nil {
		var permission system.SysPermission
		db := global.GormDB.Preload("SysRoles").Where("id = ?", id).First(&permission).Delete(&permission)
		if len(permission.SysRoles) > 0 {
			err = global.GormDB.Model(&permission).Association("SysRoles").Delete(&permission.SysRoles)
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 批量删除权限菜单
// @param: id float64
// @return: err error
func (permissionService *PermissionService) DeleteByIds(ids []uint64) (err error) {
	err = global.GormDB.Transaction(func(tx *gorm.DB) error {

		for _, id := range ids {
			err = deleteById(id, ids, tx)
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}

func deleteById(id uint64, ids []uint64, tx *gorm.DB) (err error) {

	err = tx.Where("permission_id = ?", id).First(&system.SysRolePermission{}).Error
	if err != nil {
		var permission system.SysPermission
		db := tx.Preload("SysRoles").Where("id = ?", id).First(&permission).Delete(&permission)
		if len(permission.SysRoles) > 0 {
			err = tx.Model(&permission).Association("SysRoles").Delete(&permission.SysRoles)
		} else {
			err = db.Error
		}

		var children []system.SysPermission
		err = tx.Where("parent_id = ?", id).Find(&children).Error
		if len(children) > 0 {
			for _, child := range children {

				if !utils.JudgeIds(child.ID, ids) {
					err = deleteById(child.ID, ids, tx)
				}

			}
		}

	} else {
		return errors.New("删除失败，包含正被用户使用关联的菜单")
	}

	return err
}

// Update
// @function: Update
// @description: 更新权限菜单
// @param: permission request.PermissionFormRequest
// @return: err error
func (permissionService *PermissionService) Update(permission *request.PermissionFormRequest) (err error) {
	var oldPermission system.SysPermission
	dataMap := make(map[string]interface{})

	dataMap["keep_alive"] = permission.KeepAlive
	dataMap["parent_id"] = permission.ParentId
	dataMap["path"] = permission.Path
	dataMap["name"] = permission.Name
	dataMap["hidden"] = permission.Hidden
	dataMap["component"] = permission.Component
	dataMap["title"] = permission.Title
	dataMap["sort_order"] = permission.SortOrder
	dataMap["redirect"] = permission.Redirect
	dataMap["is_button"] = permission.IsButton
	dataMap["status"] = permission.Status
	dataMap["redirect"] = permission.Redirect
	dataMap["api"] = permission.Api

	if permission.Icon == "" {
		dataMap["icon"] = permission.Icon
	} else {
		dataMap["icon"] = "menu"
	}

	if permission.Path == "" && permission.Component != "" {
		dataMap["path"] = permission.Name
	}

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", permission.Id).Find(&oldPermission)
		if oldPermission.Name != permission.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", permission.Id, permission.Name).First(&system.SysPermission{}).Error, gorm.ErrRecordNotFound) {
				global.Logger.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&system.SysPermission{}, "parent_id = ? and is_button=1", permission.Id).Error
		if txErr != nil {
			global.Debug("UpdatePermission", txErr.Error())
			return txErr
		}
		var buttons []system.SysPermission
		if len(permission.Buttons) > 0 {
			for _, button := range permission.Buttons {
				if button != "" {
					s := strings.Split(button, ":")
					if len(s) >= 2 {
						p := system.SysPermission{}

						p.ParentId = oldPermission.ID
						p.Name = s[0]
						p.Title = s[1]
						p.Icon = "plus-circle"

						buttons = append(buttons, p)
					}
				}

			}
			txErr = tx.Create(&buttons).Error
			if txErr != nil {
				global.Debug("UpdatePermission", txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(dataMap).Error
		if txErr != nil {
			global.Debug("UpdatePermission", txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

// GetById
// @function: GetById
// @description: 返回当前选中permission
// @param: id float64
// @return: err error, permission model.SysPermission
func (permissionService *PermissionService) GetById(id float64) (err error, permission system.SysPermission) {
	err = global.GormDB.Where("id = ?", id).First(&permission).Error
	return
}

// GetList
// @function: GetList
// @param: permission request.SearchPermissionParams
// @description: 获取权限列表
// @return: err error, list interface{}, total int64
func (permissionService *PermissionService) GetList(params request.SearchPermissionParams) (err error, list interface{}, total int64) {

	var allPermissions []system.SysPermission
	db := global.GormDB.Order("sort_order asc").Order("id asc")

	if params.Title != "" {
		db.Where("title = '%?%' ", params.Title)
	}

	err = db.Find(&allPermissions).Error

	len := len(allPermissions)

	return err, allPermissions, int64(len)
}

// GetTree
// @function: GetTree
// @param: permission request.SearchPermissionParams
// @description: 获取基础路由树
// @return: err error, menus []model.SysPermission
func (permissionService *PermissionService) GetTree(params request.SearchPermissionParams) (err error, menus []system.SysPermission) {
	err, treeMap := permissionService.getTreeMap(params)
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = permissionService.getChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

// @function: getTreeMap
// @param: permission request.SearchPermissionParams
// @description: 获取路由总树map
// @return: err error, treeMap map[string][]model.SysBasePermissionMenu
func (permissionService *PermissionService) getTreeMap(params request.SearchPermissionParams) (err error, treeMap map[uint64][]system.SysPermission) {
	var allPermissions []system.SysPermission
	treeMap = make(map[uint64][]system.SysPermission)

	db := global.GormDB.Order("sort_order asc").Order("id asc")

	if params.Title != "" {
		db.Where("title = '%?%' ", params.Title)
	}

	err = db.Find(&allPermissions).Error

	for _, v := range allPermissions {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// @function: getPermissionChildrenList
// @description: 获取菜单的子菜单
// @param: menu *model.SysBasePermissionMenu, treeMap map[string][]model.SysBasePermissionMenu
// @return: err error
func (permissionService *PermissionService) getChildrenList(item *system.SysPermission, treeMap map[uint64][]system.SysPermission) (err error) {
	item.Children = treeMap[item.ID]
	for i := 0; i < len(item.Children); i++ {
		err = permissionService.getChildrenList(&item.Children[i], treeMap)
	}
	return err
}

// GetByRoleId
// @function: GetByRoleId
// @param: roleId uint
// @description: 获取权限列表
// @return: err error, list interface{}, total int64
func (permissionService *PermissionService) GetByRoleId(roleId uint) (err error, list interface{}) {

	var rolePermissions []system.SysRolePermission
	err = global.GormDB.Where("role_id = ?", roleId).Find(&rolePermissions).Error

	return err, rolePermissions
}
