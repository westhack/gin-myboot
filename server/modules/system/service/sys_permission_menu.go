package service

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"
)

type PermissionMenuService struct {
}

var PermissionMenuAppService = new(PermissionMenuService)

// GetUserRolePermissionList
// @function: GetUserRolePermissionMenuList
// @description: 获取用户全部菜单权限列表
// @param: roleId string
// @return: err error, menus []system.SysPermission
func (permissionMenuService *PermissionMenuService) GetUserRolePermissionList(roleId uint64) (err error, userPermissions []system.SysPermission) {

	var rolePermissionMenus []system.SysRolePermission
	var permissionIds []uint64
	err = global.GormDB.Where("role_id = ?", roleId).Find(&rolePermissionMenus).Error
	for _, v := range rolePermissionMenus {
		permissionIds = append(permissionIds, v.PermissionId)
	}

	err = global.GormDB.Where("id in ?", permissionIds).Where("status = 1").Order("sort_order asc").Order("id asc").Find(&userPermissions).Error

	return err, userPermissions
}

func (permissionMenuService *PermissionMenuService) GetPermissionMenuTreeMap(roleId uint64) (err error, treeMap map[uint64][]system.SysRolePermissionMenu) {
	var allPermissionMenus []system.SysRolePermissionMenu

	treeMap = make(map[uint64][]system.SysRolePermissionMenu)

	err = global.GormDB.Where("role_id = ?", roleId).
		Where("is_button = 0").
		Order("sort_order").Find(&allPermissionMenus).Error

	for _, v := range allPermissionMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}


//@function: GetRolePermissionMenuTree
//@description: 获取用户角色动态权限菜单树
//@param: roleId string
//@return: err error, menus []model.SysRolePermissionMenu

func (permissionMenuService *PermissionMenuService) GetUserRolePermissionMenuTree(roleId uint64) (err error, menus []system.SysRolePermissionMenu) {
	err, menuTree := permissionMenuService.GetPermissionMenuTreeMap(roleId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = permissionMenuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}


// @function: getChildrenList
// @description: 获取子菜单
// @param: menu *model.SysRolePermissionMenu, treeMap map[string][]model.SysRolePermissionMenu
// @return: err error
func (permissionMenuService *PermissionMenuService) getChildrenList(menu *system.SysRolePermissionMenu, treeMap map[uint64][]system.SysRolePermissionMenu) (err error) {
	menu.Children = treeMap[menu.PermissionId]
	for i := 0; i < len(menu.Children); i++ {
		err = permissionMenuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// AddPermissionMenuRole
// @function: AddPermissionMenuRole
// @description: 为角色增加menu树
// @param: menus []model.SysBasePermissionMenu, roleId string
// @return: err error
func (permissionMenuService *PermissionMenuService) AddPermissionMenuRole(menus []system.SysPermission, roleId uint64) (err error) {
	var role system.SysRole
	role.Permissions = menus
	err = RoleServiceApp.SetPermissionMenuRole(&role)
	return err
}

// GetPermissionMenuRole
// @function: GetPermissionMenuRole
// @description: 查看当前角色树
// @param: info *request.GetRoleId
// @return: err error, menus []model.SysRolePermissionMenu
func (permissionMenuService *PermissionMenuService) GetPermissionMenuRole(info *request.GetRoleId) (err error, menus []system.SysRolePermissionMenu) {
	err = global.GormDB.Where("role_id = ? ", info.RoleId).Order("sort_order").Find(&menus).Error
	//sql := "SELECT role_menu.keep_alive,role_menu.default_menu,role_menu.created_at,role_menu.updated_at,role_menu.deleted_at,role_menu.menu_level,role_menu.parent_id,role_menu.path,role_menu.`name`,role_menu.hidden,role_menu.component,role_menu.title,role_menu.icon,role_menu.sort,role_menu.menu_id,role_menu.role_id FROM role_menu WHERE role_menu.role_id = ? ORDER BY role_menu.sort ASC"
	//err = global.GormDB.Raw(sql, roleId).Scan(&menus).Error
	return err, menus
}
