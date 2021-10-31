package service

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	uuid "github.com/satori/go.uuid"
)

type MessageService struct {
}

// Create
// @function: Create
// @description: 创建消息
// @param: sysMessage *request.MessageFormRequest
// @return: err error
func (messageService *MessageService) Create(sysMessage *systemReq.MessageFormRequest) (err error) {
	var messages []system.SysMessage
	if len(sysMessage.UserID) == 0 {
		return errors.New("接受用户不存在")
	}
	for _, userId := range sysMessage.UserID {
		messages = append(messages, system.SysMessage{
			UserId:     userId,
			FromUserId: sysMessage.FormUserID,
			Type:       sysMessage.Type,
			Title:      sysMessage.Title,
			Content:    sysMessage.Content,
			Icon:       sysMessage.Icon,
			Status:     false,
			SessionId:  uuid.NewV4().String(),
		})
	}

	err = global.GormDB.Create(&messages).Error

	return err
}

// Update
// @function: Update
// @description: 修改消息
// @param: sysMessage system.SysMessage
// @return: err error
func (messageService *MessageService) Update(sysMessage system.SysMessage) (err error) {
	err = global.GormDB.Updates(&sysMessage).Error
	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 批量删除记录
// @param: ids []uint64
// @return: err error
func (messageService *MessageService) DeleteByIds(ids []uint64) (err error) {
	err = global.GormDB.Delete(&[]system.SysMessage{}, "id in (?)", ids).Error
	return err
}

// Delete
// @function: Delete
// @description: 删除操作记录
// @param: id uint64
// @return: err error
func (messageService *MessageService) Delete(id uint64) (err error) {
	err = global.GormDB.Delete(&[]system.SysMessage{}, "id = ?", id).Error
	return err
}

// GetById
// @function: GetById
// @description: 根据id获取单条操作记录
// @param: id uint
// @return: err error, sysMessage model.SysMessage
func (messageService *MessageService) GetById(id uint64) (err error, sysMessage system.SysMessage) {
	err = global.GormDB.Where("id = ?", id).First(&sysMessage).Error
	return err, sysMessage
}

// SetUserMessageView
// @function: SetUserMessageView
// @description: 根据id获取单条操作记录
// @param: id uint
// @return: err error
func (messageService *MessageService) SetUserMessageView(id uint64, userId uint64) (err error) {

	var message system.SysMessage

	data := make(map[string]interface{})
	data["view_time"] = utils.GetDateTime()
	data["status"] = 1

	err = global.GormDB.Model(&message).Where("id = ? and user_id = ?", id, userId).Updates(data).Error

	if err != nil {
		global.Error("==>", err)
	}
	return err
}


// UserDelete
// @function: UserDelete
// @description: 用户删除消息
// @param: id uint64
// @param: userId uint64
// @return: err error
func (messageService *MessageService) UserDelete(id uint64, userId uint64) (err error) {

	err, message := messageService.GetById(id)
	if err != nil {
		return errors.New("消息不存在")
	}

	if message.UserId != userId {
		return errors.New("消息不存在")
	}

	// 删除
	err = global.GormDB.Delete(&system.SysMessage{}, "id = ? and user_id = ?", id, userId).Error

	return err
}

// CountByUserIdAndStatus
// @function: CountByUserIdAndStatus
// @description: 根据用户ID和状态统计数量
// @param: id uint
// @return: err error, sysMessage model.SysMessage
func (messageService *MessageService) CountByUserIdAndStatus(userId uint64, status int) (count int64) {
	db := global.GormDB.Model(&system.SysMessage{}).Where("user_id", userId)
	if status == 0 {
		db.Where("status", 0)
	} else if status == 1 {
		db.Where("status", 1)
	}

	_ = db.Count(&count).Error
	return count
}

// SelectPageByUserIdAndStatus
// @function: SelectPageByUserIdAndStatus
// @description: 分页用户消息列表
// @param: info queryParams request.QueryParams
// @return: err error, list interface{}, total int64
func (messageService *MessageService) SelectPageByUserIdAndStatus(queryParams request.PageInfo, userId uint64, status int) (err error, res interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)

	db := global.GormDB.Model(&system.SysMessage{}).Preload("FormUser").Where("user_id", userId)
	if status == 0 {
		db.Where("status", 0)
	} else if status == 1 {
		db.Where("status", 1)
	}

	var list []system.SysMessage

	err = db.Count(&total).Error

	if err != nil {
		return err, list, total
	} else {
		err = db.Limit(limit).Offset(offset).Order("created_At DESC").Find(&list).Error

		formUser := system.SysUser{}
		formUser.ID = 0
		formUser.Username = "系统"
		for i, message := range list {
			if message.FormUser.ID == 0 {
				message.FormUser = formUser
			}
			list[i] = message
		}
	}

	return err, list, total
}

// GetList
// @function: GetList
// @description: 分页获取操作记录列表
// @param: info queryParams request.QueryParams
// @return: err error, list interface{}, total int64
func (messageService *MessageService) GetList(queryParams request.QueryParams) (err error, res interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)
	// 创建db
	db := global.GormDB.Model(&system.SysMessage{}).Preload("FormUser").Preload("User")
	var list []system.SysMessage
	db.Scopes(model.Search(queryParams.Search))

	err = db.Count(&total).Error

	if err != nil {
		return err, list, total
	} else {

		if queryParams.SortOrder.Column == "" {
			queryParams.SortOrder.Column = "id"
			queryParams.SortOrder.Order = "desc"
		}

		err = db.Scopes(model.SortOrder(queryParams.SortOrder)).Limit(limit).Offset(offset).Find(&list).Error

		formUser := system.SysUser{}
		formUser.ID = 0
		formUser.Username = "系统"

		for i, message := range list {
			if message.FormUser.ID == 0 {
				message.FormUser = formUser
			}
			list[i] = message
		}

	}

	return err, list, total
}
