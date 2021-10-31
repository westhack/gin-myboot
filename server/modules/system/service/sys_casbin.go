package service

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/system/model/request"
	system "gin-myboot/modules/system/model"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"sync"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

// UpdateCasbin
// @function: UpdateCasbin
// @description: 更新casbin权限
// @param: roleId string, casbinList []request.CasbinInfo
// @return: error
func (casbinService *CasbinService) UpdateCasbin(roleId uint64, casbinInfos []request.CasbinInfo) error {
	roleIdStr := strconv.FormatUint(roleId, 10)
	casbinService.ClearCasbin(0, roleIdStr)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := system.CasbinModel{
			Ptype:       "p",
			RoleId:      roleIdStr,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.RoleId, cm.Path, cm.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// UpdateCasbinApi
// @function: UpdateCasbinApi
// @description: API更新随动
// @param: oldPath string, newPath string, oldMethod string, newMethod string
// @return: error
func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GormDB.Table("casbin_rule").Model(&system.CasbinModel{}).
		Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetPolicyPathByRoleId
// @function: GetPolicyPathByRoleId
// @description: 获取权限列表
// @param: roleId string
// @return: pathMaps []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByRoleId(roleId uint64) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	roleIdStr := strconv.Itoa(int(roleId))
	list := e.GetFilteredPolicy(0, roleIdStr)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin
// @function: ClearCasbin
// @description: 清除匹配的权限
// @param: v int, p ...string
// @return: bool
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Casbin
// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.Enforcer
func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GormDB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.Config.Casbin.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", casbinService.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

// ParamsMatch
// @function: ParamsMatch
// @description: 自定义规则函数
// @param: fullNameKey1 string, key2 string
// @return: bool
func (casbinService *CasbinService) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc
// @function: ParamsMatchFunc
// @description: 自定义规则函数
// @param: args ...interface{}
// @return: interface{}, error
func (casbinService *CasbinService) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return casbinService.ParamsMatch(name1, name2), nil
}
