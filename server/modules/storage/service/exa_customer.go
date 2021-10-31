package service

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	storage "gin-myboot/modules/storage/model"
	system "gin-myboot/modules/system/model"
	systemService "gin-myboot/modules/system/service"
)

type CustomerService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) CreateExaCustomer(e storage.ExaCustomer) (err error) {
	err = global.GormDB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) DeleteExaCustomer(e storage.ExaCustomer) (err error) {
	err = global.GormDB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (exa *CustomerService) UpdateExaCustomer(e *storage.ExaCustomer) (err error) {
	err = global.GormDB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: err error, customer model.ExaCustomer

func (exa *CustomerService) GetExaCustomer(id uint64) (err error, customer storage.ExaCustomer) {
	err = global.GormDB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserRoleID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func (exa *CustomerService) GetCustomerInfoList(sysUserRoleID uint64, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GormDB.Model(&storage.ExaCustomer{})
	var a system.SysRole
	a.ID = sysUserRoleID
	err, role := systemService.RoleServiceApp.GetRoleInfo(a)
	var dataId []uint64
	for _, v := range role.RoleDepartments {
		dataId = append(dataId, v.RoleId)
	}
	var CustomerList []storage.ExaCustomer
	err = db.Where("sys_user_role_id in ?", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_role_id in ?", dataId).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
