package service

import (
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"gorm.io/gorm"
)

type DepartmentService struct {
}

// Create
// @function: Create
// @description: 更新
// @param: department model.SysDepartment
// @return: err error
func (departmentService *DepartmentService) Create(department *system.SysDepartment) (err error) {

	err = global.GormDB.Create(department).Error
	if err == nil {
		departmentService.saveHeader(department, global.GormDB)
	}

	return err
}

func (departmentService *DepartmentService) saveHeader(department *system.SysDepartment, tx *gorm.DB)  {
	if len(department.MainHeader) > 0 {
		var sysDepartmentHeaders []system.SysDepartmentHeader
		for _, userId := range department.MainHeader {
			sysDepartmentHeaders = append(sysDepartmentHeaders, system.SysDepartmentHeader{
				DepartmentId: department.ID,
				UserId: userId,
				Type: 0,
			})
		}
		tx.Create(&sysDepartmentHeaders)
	}
	if len(department.ViceHeader) > 0 {
		var sysDepartmentHeaders []system.SysDepartmentHeader
		for _, userId := range department.ViceHeader {
			sysDepartmentHeaders = append(sysDepartmentHeaders, system.SysDepartmentHeader{
				DepartmentId: department.ID,
				UserId: userId,
				Type: 1,
			})
		}
		tx.Create(&sysDepartmentHeaders)
	}
}

// Update
// @function: Update
// @description: 更新
// @param: department model.SysDepartment
// @return: err error
func (departmentService *DepartmentService) Update(department *system.SysDepartment) (err error) {
	var oldDepartment system.SysDepartment

	upDateMap := make(map[string]interface{})

	if department.ID == department.ParentId {
		return errors.New("上级节点不能为自己")
	}

	upDateMap["title"] = department.Title
	upDateMap["status"] = department.Status
	upDateMap["parent_id"] = department.ParentId
	upDateMap["sort_order"] = department.SortOrder

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", department.ID).First(&oldDepartment).Error
		if err != nil {
			return err
		}

		txErr := tx.Where("id = ?", department.ID).First(&oldDepartment).Updates(upDateMap).Error
		if txErr != nil {
			global.Logger.Debug(txErr.Error())
			return txErr
		}

		err = tx.Unscoped().Where("department_id = ?", department.ID).Delete(&system.SysDepartmentHeader{}).Error
		if err != nil {
			return err
		}

		departmentService.saveHeader(department, tx)

		return nil
	})
	return err
}

// Delete
// @function: Delete
// @description: 删除部门
// @param: id float64
// @return: err error
func (departmentService *DepartmentService) Delete(id uint64) (err error) {

	global.GormDB.Transaction(func(tx *gorm.DB) error {
		err := departmentService.delete(id, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 批量删除部门
// @param: id float64
// @return: err error
func (departmentService *DepartmentService) DeleteByIds(ids []uint64) (err error) {

	global.GormDB.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			err := departmentService.delete(id, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (departmentService *DepartmentService) delete(id uint64, txDb *gorm.DB) (err error) {
	var department system.SysDepartment
	err = txDb.Where("id = ?", id).First(&department).Error
	err = txDb.Where("parent_id = ?", id).First(&system.SysDepartment{}).Error
	if err != nil {
		return errors.New(department.Title + " 存在子部门不可删除")
	}

	err = txDb.Where("department_id = ?", id).First(&system.SysUser{}).Error
	if err != nil {
		return errors.New(department.Title + " 存在子部门不可删除")
	}

	db := txDb.Unscoped().Where("id = ?", id).Delete(&system.SysDepartment{})
	if len(department.RoleDepartments) > 0 {
		err = txDb.Model(&department).Unscoped().Association("RoleDepartments").Delete(&department.RoleDepartments)
	} else {
		err = db.Error
	}

	return err
}

// GetList
// @function: GetList
// @description: 分页获取数据
// @param: info request.SearchDepartmentParams
// @return: err error, list interface{}, total int64
func (departmentService *DepartmentService) GetList(req systemReq.SearchDepartmentParams) (err error, list interface{}, total int64) {
	db := global.GormDB
	var departments []system.SysDepartment

	if req.Title != "" {
		db.Where("title = ?", req.Title)
	}
	err = db.Find(&departments).Error
	//err = db.Where("parent_id >= 0").Find(&department).Error
	//if len(department) > 0 {
	//	for k := range department {
	//		err = departmentService.findChildren(&department[k])
	//	}
	//}

	var treeNodes []utils.TreeNode
	for _, obj := range departments {
		var departmentHeaders []system.SysDepartmentHeader

		global.GormDB.Where("department_id = ? and type = 0", obj.ID).Preload("User").Find(&departmentHeaders)
		if len(departmentHeaders) > 0 {
			for _, header := range departmentHeaders {
				obj.MainHeaders = append(obj.MainHeaders, header.User)
				obj.MainHeader = append(obj.MainHeader, header.User.ID)
			}
		}

		global.GormDB.Where("department_id = ? and type = 1", obj.ID).Preload("User").Find(&departmentHeaders)
		if len(departmentHeaders) > 0 {
			for _, header := range departmentHeaders {
				obj.ViceHeaders = append(obj.ViceHeaders, header.User)
				obj.ViceHeader = append(obj.ViceHeader, header.User.ID)
			}
		}

		treeNodes = append(treeNodes, utils.StructToMapViaJson(obj))
	}

	treeNodes = utils.GetTree(treeNodes, "parentId", "id", "0")

	return err, treeNodes, total
}

// @function: findChildren
// @description: 查询子角色
// @param: role *model.SysRole
// @return: err error
func (departmentService *DepartmentService) findChildren(department *system.SysDepartment) (err error) {
	err = global.GormDB.Where("parent_id = ?", department.ID).Find(&department.Children).Error

	if len(department.Children) > 0 {
		for k := range department.Children {
			err = departmentService.findChildren(&department.Children[k])
		}
	}

	return err
}

// GetDepartmentById
// @function: GetDepartmentById
// @description: 返回当前选中department
// @param: id uint
// @return: err error, department model.SysDepartment
func (departmentService *DepartmentService) GetDepartmentById(id uint64) (err error, department system.SysDepartment) {
	err = global.GormDB.Where("id = ?", id).First(&department).Error
	return
}

// GetAll
// @function: GetAll
// @description: 获取全部可用部门
// @return: err error, list interface{}, total int64
func (departmentService *DepartmentService) GetAll() (err error, list interface{}, total int64) {
	db := global.GormDB
	var sysDepartments []system.SysDepartment

	err = db.Where("status = 1").Find(&sysDepartments).Error

	total = int64(len(sysDepartments))

	return err, sysDepartments, total
}
