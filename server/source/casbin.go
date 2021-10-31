package source

import (
	"gin-myboot/global"
	"strings"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

//@description: casbin_rule 表数据初始化
func (c *casbin) Init() error {
	global.GormDB.AutoMigrate(gormadapter.CasbinRule{})
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == 81 {
			color.Danger.Println("\n[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}

		casbinSql := `INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/permission/getAll', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/permission/getTree', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/permission/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/permission/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/permission/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/permission/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/getAll', 'GET', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/changeStatus', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/admin/resetPassword', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/role/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/role/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/role/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/role/getAll', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/role/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/role/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/getAll', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/setValue', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/config/write', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/changeStatus', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/user/resetPassword', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/detail', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/getAll', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/getRoutes', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/api/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/detail', 'GET', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/getAll', 'GET', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dict/saveDetail', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dictDetail/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dictDetail/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dictDetail/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/dictDetail/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/storage/upload/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/storage/upload/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/storage/upload/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/storage/upload/getUserFileList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/storage/upload/file', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/getUserMessages', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/getUserUnreadMessages', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/common/auth/getUserInfo', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/common/auth/getUserMenus', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/common/auth/logout', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/common/auth/changePassword', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/common/auth/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/log/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/log/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/log/detail', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/log/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/log/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/log/deleteAll', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/department/getAll', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/department/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/department/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/department/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/department/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/department/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/message/create', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/redis/getList', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/redis/find', 'GET', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/redis/delete', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/redis/deleteByIds', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/redis/update', 'POST', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '1', '/api/v1/system/redis/create', 'POST', '', '', '');`

		sqls := strings.Split(casbinSql, "\n")

		for _, sql := range sqls {
			if err := global.GormDB.Exec(sql).Error; err != nil {
				return err
			}
		}
		color.Info.Println("\n[Mysql] --> casbin_rule 表初始数据成功!")
		return nil
	})
}
