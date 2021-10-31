package router

import (
	"gin-myboot/middleware"
	"gin-myboot/modules/system/api/v1"
	"github.com/gin-gonic/gin"
)

type SystemRouter struct {
}

func (s *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {

	var systemApiGroup = v1.ApiGroupApp

	logRouter := Router.Group("v1/system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		logRouter.POST("log/create", systemApiGroup.UserLogApi.Create)
		logRouter.POST("log/delete", systemApiGroup.UserLogApi.Delete)
		logRouter.POST("log/deleteByIds", systemApiGroup.UserLogApi.DeleteByIds)
		logRouter.POST("log/find", systemApiGroup.UserLogApi.Find)
		logRouter.POST("log/getList", systemApiGroup.UserLogApi.GetList)
		logRouter.POST("log/deleteAll", systemApiGroup.UserLogApi.DeleteAll)
	}

	sysRouter := Router.Group("v1/system").Use(middleware.JWTAuth()).Use(middleware.SystemLog()).Use(middleware.CasbinHandler())
	{
		sysRouter.POST("user/getList", systemApiGroup.UserApi.GetList)             // 获取用户列表
		sysRouter.POST("user/create", systemApiGroup.UserApi.Create)               // 创建用户
		sysRouter.POST("user/update", systemApiGroup.UserApi.Update)               // 修改用户
		sysRouter.POST("user/delete", systemApiGroup.UserApi.Delete)               // 删除用户
		sysRouter.POST("user/deleteByIds", systemApiGroup.UserApi.DeleteByIds)     // 批量删除用户
		sysRouter.POST("user/changeStatus", systemApiGroup.UserApi.ChangeStatus)   // 修改状态
		sysRouter.POST("user/resetPassword", systemApiGroup.UserApi.ResetPassword) // 重置密码

		sysRouter.GET("admin/getAll", systemApiGroup.AdminApi.GetAll)                // 获取管理员列表
		sysRouter.POST("admin/getList", systemApiGroup.AdminApi.GetList)             // 获取管理员列表
		sysRouter.POST("admin/create", systemApiGroup.AdminApi.Create)               // 创建管理员
		sysRouter.POST("admin/update", systemApiGroup.AdminApi.Update)               // 修改管理员
		sysRouter.POST("admin/delete", systemApiGroup.AdminApi.Delete)               // 删除管理员
		sysRouter.POST("admin/deleteByIds", systemApiGroup.AdminApi.DeleteByIds)     // 批量删除管理员
		sysRouter.POST("admin/changeStatus", systemApiGroup.AdminApi.ChangeStatus)   // 修改状态
		sysRouter.POST("admin/resetPassword", systemApiGroup.AdminApi.ResetPassword) // 重置密码

		sysRouter.POST("permission/getAll", systemApiGroup.PermissionApi.GetAll)   // 获取全部权限列表
		sysRouter.POST("permission/getTree", systemApiGroup.PermissionApi.GetTree) // 获取权限树形列表
		sysRouter.POST("permission/update", systemApiGroup.PermissionApi.Update)
		sysRouter.POST("permission/create", systemApiGroup.PermissionApi.Create)
		sysRouter.POST("permission/delete", systemApiGroup.PermissionApi.Delete)
		sysRouter.POST("permission/deleteByIds", systemApiGroup.PermissionApi.DeleteByIds)

		sysRouter.POST("department/getAll", systemApiGroup.DepartmentApi.GetAll) // 获取全部可用部门列表
		sysRouter.POST("department/getList", systemApiGroup.DepartmentApi.GetList)
		sysRouter.POST("department/delete", systemApiGroup.DepartmentApi.Delete)
		sysRouter.POST("department/deleteByIds", systemApiGroup.DepartmentApi.DeleteByIds)
		sysRouter.POST("department/create", systemApiGroup.DepartmentApi.Create)
		sysRouter.POST("department/update", systemApiGroup.DepartmentApi.Update)

		sysRouter.POST("api/create", systemApiGroup.SystemApiApi.Create)           // 创建Api
		sysRouter.POST("api/delete", systemApiGroup.SystemApiApi.Delete)           // 删除Api
		sysRouter.POST("api/getList", systemApiGroup.SystemApiApi.GetList)         // 获取Api列表
		sysRouter.POST("api/detail", systemApiGroup.SystemApiApi.Detail)           // 获取单条Api消息
		sysRouter.POST("api/update", systemApiGroup.SystemApiApi.Update)           // 更新api
		sysRouter.POST("api/getAll", systemApiGroup.SystemApiApi.GetAll)           // 获取所有api
		sysRouter.POST("api/deleteByIds", systemApiGroup.SystemApiApi.DeleteByIds) // 删除选中api
		sysRouter.POST("api/getRoutes", systemApiGroup.SystemApiApi.GetRoutes)     // 删除选中api

		sysRouter.POST("role/create", systemApiGroup.RoleApi.Create)           // 创建角色
		sysRouter.POST("role/delete", systemApiGroup.RoleApi.Delete)           // 删除角色
		sysRouter.POST("role/deleteByIds", systemApiGroup.RoleApi.DeleteByIds) // 删除角色
		sysRouter.POST("role/update", systemApiGroup.RoleApi.Update)           // 更新角色
		sysRouter.POST("role/getList", systemApiGroup.RoleApi.GetList)         // 获取角色列表
		sysRouter.POST("role/getAll", systemApiGroup.RoleApi.GetAll)           // 获取全部角色列表
		sysRouter.POST("role/setData", systemApiGroup.RoleApi.SetDataRole)     // 设置角色资源权限

		sysRouter.POST("config/create", systemApiGroup.ConfigApi.Create)
		sysRouter.POST("config/delete", systemApiGroup.ConfigApi.Delete)
		sysRouter.POST("config/update", systemApiGroup.ConfigApi.Update)
		sysRouter.POST("config/getList", systemApiGroup.ConfigApi.GetList)
		sysRouter.POST("config/getAll", systemApiGroup.ConfigApi.GetAll)
		sysRouter.POST("config/setValue", systemApiGroup.ConfigApi.SetValue)
		sysRouter.POST("config/write", systemApiGroup.ConfigApi.Write) // 配置写入到文件

		sysRouter.POST("dict/create", systemApiGroup.DictApi.Create)   // 新建SysDict
		sysRouter.POST("dict/delete", systemApiGroup.DictApi.Delete)   // 删除SysDict
		sysRouter.POST("dict/update", systemApiGroup.DictApi.Update)   // 更新SysDict
		sysRouter.GET("dict/find", systemApiGroup.DictApi.Find)    // 根据ID获取SysDict
		sysRouter.POST("dict/getList", systemApiGroup.DictApi.GetList) // 获取SysDict列表
		sysRouter.GET("dict/getAll", systemApiGroup.DictApi.GetAll)
		sysRouter.POST("dict/saveDetail", systemApiGroup.DictApi.SaveDetail)

		sysRouter.POST("dictDetail/create", systemApiGroup.DictDetailApi.Create)   // 新建SysDictDetail
		sysRouter.POST("dictDetail/delete", systemApiGroup.DictDetailApi.Delete)   // 删除SysDictDetail
		sysRouter.POST("dictDetail/update", systemApiGroup.DictDetailApi.Update)   // 更新SysDictDetail
		sysRouter.GET("dictDetail/find", systemApiGroup.DictDetailApi.Find)    // 根据ID获取SysDictDetail
		sysRouter.POST("dictDetail/getList", systemApiGroup.DictDetailApi.GetList) // 获取SysDictDetail列表

		sysRouter.POST("message/getList", systemApiGroup.MessageApi.GetList)
		sysRouter.POST("message/delete", systemApiGroup.MessageApi.Delete)
		sysRouter.POST("message/deleteByIds", systemApiGroup.MessageApi.DeleteByIds)
		sysRouter.POST("message/create", systemApiGroup.MessageApi.Create)
		sysRouter.POST("message/update", systemApiGroup.MessageApi.Update)
		sysRouter.POST("message/getUserMessages", systemApiGroup.MessageApi.GetUserMessages)
		sysRouter.POST("message/getUserUnreadMessages", systemApiGroup.MessageApi.GetUserUnreadMessages)
		sysRouter.POST("message/userView", systemApiGroup.MessageApi.UserView)
		sysRouter.POST("message/userDelete", systemApiGroup.MessageApi.UserDelete)

		sysRouter.POST("casbin/updateCasbin", systemApiGroup.CasbinApi.UpdateCasbin)                   // 更新角色api权限
		sysRouter.POST("casbin/getPolicyPathByRoleId", systemApiGroup.CasbinApi.GetPolicyPathByRoleId) // 获取权限列表

		sysRouter.POST("getSystemConfig", systemApiGroup.SystemApi.GetSystemConfig) // 获取配置文件内容
		sysRouter.POST("setSystemConfig", systemApiGroup.SystemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("getServerInfo", systemApiGroup.SystemApi.GetServerInfo)     // 获取服务器信息
		sysRouter.POST("reloadSystem", systemApiGroup.SystemApi.ReloadSystem)       // 重启服务

	}

}
