package service

import (
	"errors"
	"fmt"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

type ExcelService struct {
}

func (exa *ExcelService) ParseInfoList2Excel(infoList []system.SysPermission, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"})
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			menu.ID,
			menu.Name,
			menu.Path,
			menu.Hidden,
			menu.ParentId,
			menu.SortOrder,
			menu.Component,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}

func (exa *ExcelService) ParseExcel2InfoList() ([]system.SysPermission, error) {
	skipHeader := true
	fixedHeader := []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}
	file, err := excelize.OpenFile(global.Config.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return nil, err
	}
	menus := make([]system.SysPermission, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		parentId, _ := strconv.Atoi(row[4])
		hidden, _ := strconv.ParseBool(row[3])
		sort, _ := strconv.ParseFloat(row[5], 64)
		menu := system.SysPermission{
			Model: global.Model{
				ID: uint64(id),
			},
			Name:      row[1],
			Path:      row[2],
			Hidden:    hidden,
			ParentId:  uint64(parentId),
			SortOrder: sort,
			Component: row[6],
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func (exa *ExcelService) compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
