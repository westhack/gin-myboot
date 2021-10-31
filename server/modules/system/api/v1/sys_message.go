package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MessageApi struct {
}

// Create
// @Tags Message
// @Summary 创建消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MessageFormRequest true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/message/create [post]
func (a *MessageApi) Create(c *gin.Context) {
	var message systemReq.MessageFormRequest
	if err := c.ShouldBindJSON(&message); err != nil {
		response.FailWithMessage("创建失败，"+utils.GetError(err, message), c)
		return
	}

	if err := messageService.Create(&message); err != nil {
		global.Error("创建失败!", err)
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Update
// @Tags Message
// @Summary 创建消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysMessage true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/message/update [post]
func (a *MessageApi) Update(c *gin.Context) {
	var message system.SysMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		response.FailWithMessage("创建失败，"+utils.GetError(err, message), c)
		return
	}

	if err := messageService.Update(message); err != nil {
		global.Error("更新失败!", err)
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Delete
// @Tags Message
// @Summary 删除SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/message/delete [post]
func (s *MessageApi) Delete(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)
	if err := messageService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags Message
// @Summary 批量删除SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /v1/system/message/deleteByIds [post]
func (s *MessageApi) DeleteByIds(c *gin.Context) {
	var ids request.GetByIds
	_ = c.ShouldBindJSON(&ids)
	if err := messageService.DeleteByIds(ids.ID); err != nil {
		global.Logger.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// GetList
// @Tags Message
// @Summary 分页获取消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.QueryParams true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/message/getList [post]
func (s *MessageApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := messageService.GetList(searchParams); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchParams.Page,
			PageSize: searchParams.PageSize,
		}, "获取成功", c)
	}
}

// GetUserMessages
// @Tags Message
// @Summary 分页获用户消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.QueryParams true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/message/getUserMessages [post]
func (s *MessageApi) GetUserMessages(c *gin.Context) {
	var searchParams request.PageInfo
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		searchParams.PageSize = 10
		searchParams.Page = 1
	}

	err, user := utils.GetCurrUser(c)
	if err != nil {
		return
	}

	if err, list, total := messageService.SelectPageByUserIdAndStatus(searchParams, user.ID, 2); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchParams.Page,
			PageSize: searchParams.PageSize,
		}, "获取成功", c)
	}
}

// GetUserUnreadMessages
// @Tags Message
// @Summary 分页获用户未读消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.QueryParams true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/message/getUserUnreadMessages [post]
func (s *MessageApi) GetUserUnreadMessages(c *gin.Context) {
	var searchParams request.PageInfo
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		searchParams.PageSize = 10
		searchParams.Page = 1
	}

	err, user := utils.GetCurrUser(c)
	if err != nil {
		return
	}

	if err, list, total := messageService.SelectPageByUserIdAndStatus(searchParams, user.ID, 0); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {

		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchParams.Page,
			PageSize: searchParams.PageSize,
		}, "获取成功", c)
	}
}

// UserView
// @Tags Message
// @Summary 用户查看消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v1/system/message/userView [post]
func (s *MessageApi) UserView(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)
	if err := utils.Verify(id, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, user := utils.GetCurrUser(c)
	if err != nil {
		return
	}

	if err, res := messageService.GetById(id.ID); err != nil && res.UserId == user.ID {
		global.Logger.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败", c)
	} else {
		messageService.SetUserMessageView(id.ID, user.ID)
		total := messageService.CountByUserIdAndStatus(user.ID, 0)
		response.OkWithDetailed(gin.H{"message": res, "total": total}, "操作成功", c)
	}
}

// UserDelete
// @Tags Message
// @Summary 用户删除消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/message/userDelete [post]
func (s *MessageApi) UserDelete(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)
	if err := utils.Verify(id, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, user := utils.GetCurrUser(c)
	if err != nil {
		return
	}

	if err := messageService.UserDelete(id.ID, user.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败，"+err.Error(), c)
	} else {
		total := messageService.CountByUserIdAndStatus(user.ID, 0)
		response.OkWithDetailed(gin.H{"total": total}, "删除成功", c)
	}
}
