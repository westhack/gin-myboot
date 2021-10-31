package service

import (
	"encoding/json"
	common "gin-myboot/modules/common/model"
	"io/ioutil"
)

type ModuleService struct {
}

// GetModules
// @function: GetModules
// @description: 获取模块列表
// @return: err error, modules []common.Module
func (moduleService *ModuleService) GetModules() (err error, modules []common.Module) {
	files, err := ioutil.ReadDir("modules")
	for _, fi := range files {
		if fi.IsDir() {
			fileName := "modules/" + fi.Name() + "/module.json"
			file, err := ioutil.ReadFile(fileName)
			module := common.Module{}
			if err == nil {
				err := json.Unmarshal(file, &module)
				if err == nil {
					module.Name = fi.Name()
					if module.Title == "" {
						module.Title = module.Name
					}
					modules = append(modules, module)
				}
			} else {
				module.Name = fi.Name()
				module.Title = module.Name
				modules = append(modules, module)
				var b []byte
				b, err := json.Marshal(&module)
				if err != nil {
					continue
				}

				_ = ioutil.WriteFile(fileName, b, 0775)
			}
		}
	}

	return err, modules
}
