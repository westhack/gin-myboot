package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
)

var RolePermissionMenu = new(rolePermissionMenu)

type rolePermissionMenu struct{}

//@description: role_permission_menus 视图数据初始化
func (a *rolePermissionMenu) Init() error {
	if global.GormDB.Find(&[]system.SysRolePermissionMenu{}).RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> role_permission_menus 视图已存在!")
		return nil
	}
	if err := global.GormDB.Exec("CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `role_permission_menus` AS select `sys_permissions`.`id` AS `id`,`sys_permissions`.`created_at` AS `created_at`,`sys_permissions`.`updated_at` AS `updated_at`,`sys_permissions`.`deleted_at` AS `deleted_at`,`sys_permissions`.`level` AS `level`,`sys_permissions`.`parent_id` AS `parent_id`,`sys_permissions`.`path` AS `path`,`sys_permissions`.`name` AS `name`,`sys_permissions`.`hidden` AS `hidden`,`sys_permissions`.`component` AS `component`,`sys_permissions`.`title` AS `title`,`sys_permissions`.`icon` AS `icon`,`sys_role_permissions`.`role_id` AS `role_id`,`sys_role_permissions`.`permission_id` AS `permission_id`,`sys_permissions`.`keep_alive` AS `keep_alive`,`sys_permissions`.`status` AS `status`,`sys_permissions`.`is_button` AS `is_button`,`sys_permissions`.`redirect` AS `redirect`,`sys_permissions`.`api` AS `api`,`sys_permissions`.`sort_order` AS `sort_order` from (`sys_role_permissions` join `sys_permissions` on((`sys_role_permissions`.`permission_id` = `sys_permissions`.`id`)))").Error; err != nil {
		return err
	}
	color.Info.Println("\n[Mysql] --> role_permission_menus 视图创建成功!")
	return nil
}
