package service

import (
	"encoding/json"
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"gorm.io/gorm"
)

type ConfigService struct{}

// Create
// @function: Create
// @description: 创建字典数据
// @param: sysConfig model.Config
// @return: err error
func (configService *ConfigService) Create(sysConfig system.SysConfig) (err error) {
	if (!errors.Is(global.GormDB.First(&system.SysConfig{}, "name = ?", sysConfig.Name).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的 name，不允许创建")
	}
	err = global.GormDB.Create(&sysConfig).Error
	return err
}

// Delete
// @function: Delete
// @description: 删除字典数据
// @param: sysConfig model.Config
// @return: err error
func (configService *ConfigService) Delete(sysConfig system.SysConfig) (err error) {
	err = global.GormDB.Delete(&sysConfig).Error
	return err
}

// Update
// @function: Update
// @description: 更新字典数据
// @param: sysConfig *model.Config
// @return: err error
func (configService *ConfigService) Update(sysConfig *system.SysConfig) (err error) {
	var config system.SysConfig
	sysConfigMap := map[string]interface{}{
		"parent_name": sysConfig.ParentName,
		"label":       sysConfig.Label,
		"name":        sysConfig.Name,
		"type":        sysConfig.Type,
		"status":      sysConfig.Status,
		"data_source": sysConfig.DataSource,
		"rule_source": sysConfig.RuleSource,
		"description": sysConfig.Description,
		"multiple":    sysConfig.Multiple,
		"sort_order":  sysConfig.SortOrder,
		"module":      sysConfig.Module,
		"value_type":  sysConfig.ValueType,
	}
	db := global.GormDB.Where("id = ?", sysConfig.ID).First(&config)
	if config.Name == sysConfig.Name {
		err = db.Updates(sysConfigMap).Error
	} else {
		if (!errors.Is(global.GormDB.First(&system.SysConfig{}, "name = ?", sysConfig.Name).Error, gorm.ErrRecordNotFound)) {
			return errors.New("存在相同的 name，不允许创建")
		}
		err = db.Updates(sysConfigMap).Error

	}
	return err
}

// GetConfig
// @function: GetConfig
// @description: 根据id或者type获取字典单条数据
// @param: Type string, Id uint
// @return: err error, sysConfig model.Config
func (configService *ConfigService) GetConfig(Type string, Id uint64) (err error, sysConfig system.SysConfig) {
	err = global.GormDB.Where("type = ? OR id = ?", Type, Id).Preload("ConfigDetails").First(&sysConfig).Error
	return
}

// GetList
// @function: GetList
// @description: 分页获取列表
// @param: info request.ConfigSearch
// @return: err error, list interface{}, total int64
func (configService *ConfigService) GetList(req request.ConfigSearch) (err error, list interface{}, total int64) {

	db := global.GormDB.Model(&system.SysConfig{}).Where("status=1")
	var sysConfigs []system.SysConfig

	err = db.Count(&total).Error
	err = db.Order("sort_order ASC").Find(&sysConfigs).Error
	return err, sysConfigs, total
}

// GetAll
// @function: GetAll
// @description: 获取全部数据
// @param: info request.ConfigSearch
// @return: err error, list interface{}, total int64
func (configService *ConfigService) GetAll(req request.ConfigSearch) (err error, list interface{}, total int64) {

	db := global.GormDB.Model(&system.SysConfig{})
	var sysConfigs []system.SysConfig

	err = db.Count(&total).Error
	err = db.Order("sort_order ASC").Find(&sysConfigs).Error
	return err, sysConfigs, total
}

// SetValue
// @function: SetValue
// @description: 更新配置值
// @param: sysConfig *model.Config
// @return: err error
func (configService *ConfigService) SetValue(values map[string]string) (err error) {

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		for name, value := range values {
			//db := tx.Where("name = ?", name).First(&config)

			//err = db.Update("value", value).Error
			err = tx.Model(&system.SysConfig{}).Where("name = ?", name).Update("value", value).Error

			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// Write
// @description   set system config,
// @function: configWrite
// @description: 设置配置文件
// @param: system model.System
// @return: err error
func (configService *ConfigService) Write(configs map[string]interface{}) (err error) {
	var sysConfigs []system.SysConfig
	err = global.GormDB.Where("parent_name != ''").Find(&sysConfigs).Error
	if err != nil {
		return err
	}

	cs := make(map[string]interface{})
	for _, config := range sysConfigs {
		if config.Type == "single-dynamic-input" {
			var v map[string]interface{}
			err = json.Unmarshal([]byte(config.Value), &v)
			if err != nil {
				global.Error("====> json Unmarshal"+config.Name, config.Value)
			}
			global.Error("====> json "+config.Name, v)
			cs[config.Name] = v
		} else if config.Type == "dynamic-input" || config.Type == "dict" || config.Type == "banner" {
			var v []map[string]interface{}
			err = json.Unmarshal([]byte(config.Value), &v)
			if err != nil {
				global.Error("====> json Unmarshal"+config.Name, config.Value)
			}
			global.Error("====> json "+config.Name, v)
			cs[config.Name] = v
		} else {
			cs[config.Name] = config.Value
		}
	}

	global.Error("json ", cs)

	//cs := utils.StructToMap(config)
	for k, v := range cs {
		global.Viper.Set(k, v)
	}
	err = global.Viper.WriteConfig()
	return err
}
