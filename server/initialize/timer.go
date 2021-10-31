package initialize

import (
	"fmt"
	"gin-myboot/config"
	"gin-myboot/global"
	"gin-myboot/utils"
)

func Timer() {
	if global.Config.Timer.Start {
		for _, detail := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("ClearDB", global.Config.Timer.Spec, func() {
					err := utils.ClearTable(global.GormDB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
