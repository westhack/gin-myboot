package service

import (
	"database/sql"
	"fmt"
	"gin-myboot/config"
	"gin-myboot/global"
	generator "gin-myboot/modules/generator/model"
	"gin-myboot/modules/install/model"
	storage "gin-myboot/modules/storage/model"
	system "gin-myboot/modules/system/model"
	"gin-myboot/source"
	"gin-myboot/utils"
	"path/filepath"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InitDBService struct {
}

var InitDBServiceApp = new(InitDBService)


// @function: writeConfig
// @description: 回写配置
// @param: viper *viper.Viper, mysql config.Mysql
// @return: error
func (initDBService *InitDBService) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.Config.Mysql = mysql
	cs := utils.StructToMap(global.Config)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

// @function: createTable
// @description: 创建数据库(mysql)
// @param: dsn string, driver string, createSql
// @return: error
func (initDBService *InitDBService) createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func (initDBService *InitDBService) initDB(InitDBFunctions ...model.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// InitDB
// @function: InitDB
// @description: 创建数据库并初始化
// @param: conf model.Mysql
// @return: error
func (initDBService *InitDBService) InitDB(conf model.Mysql) error {

	if conf.Path == "" {
		conf.Path = "127.0.0.1:3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", conf.Username, conf.Password, conf.Path)
	global.Error("=======>", dsn)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.Dbname)
	if err := initDBService.createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     conf.Path,
		Dbname:   conf.Dbname,
		Username: conf.Username,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if MysqlConfig.Dbname == "" {
		return nil
	}

	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
	global.Error("=======>", linkDns)
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.GormDB = db
	}

	err := global.GormDB.AutoMigrate(
		system.SysConfig{},
		system.SysUser{},
		system.SysRole{},
		system.SysUserRole{},
		system.SysApi{},
		system.SysPermission{},
		system.SysRolePermission{},
		system.JwtBlacklist{},
		system.SysDict{},
		system.SysDictDetail{},
		system.SysLog{},
		system.SysDepartment{},
		system.SysDepartmentHeader{},
		system.SysRoleDepartment{},
		system.SysMessage{},
		system.SysDict{},
		system.SysDictDetail{},
		storage.SysUploadFile{},
		storage.ExaFile{},
		storage.ExaFileChunk{},
		storage.ExaSimpleUploader{},
		storage.ExaCustomer{},
		generator.SysAutoCodeHistory{},
	)
	if err != nil {
		global.GormDB = nil
		return err
	}
	err = initDBService.initDB(
		source.Admin,
		source.Role,
		source.Permission,
		source.RolePermission,
		source.Casbin,
		source.Dict,
		source.DictDetail,
		source.File,
		source.UserRole,
		source.Department,
		source.Config,
		source.RolePermissionMenu,
	)
	if err != nil {
		global.GormDB = nil
		return err
	}
	if err = initDBService.writeConfig(global.Viper, MysqlConfig); err != nil {
		return err
	}
	global.Config.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}
