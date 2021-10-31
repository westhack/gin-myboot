package api

import (
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	common "gin-myboot/modules/common/service"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	SortOrder     request.SortOrderParams
	IdName        string
	StatusColumn  string
	StatusEnable  uint
	StatusDisable uint
	Service       common.BaseServiceInterface
	Rules         Rules
}

type Rules struct {
	Create utils.Rules
	Update utils.Rules
}

func (api *BaseApi) SetService(service common.BaseServiceInterface) {
	api.Service = service
}

func (api *BaseApi) GetService() (ret common.BaseServiceInterface) {
	return api.Service
}

func (api *BaseApi) GetAll(c *gin.Context) {
	service := api.GetService()
	err, ret := service.GetAll()

	if err != nil {
		response.FailWithMessage("操作失败"+err.Error(), c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	err, ret := service.GetList(searchParams)

	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Create(c *gin.Context) {

	var req map[string]interface{}
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, api.Rules.Create); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	err, ret := service.Create(req)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Update(c *gin.Context) {
	var req map[string]interface{}
	_ = c.ShouldBindJSON(&req)
	if err := utils.VerifyMap(req, api.Rules.Update); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	err, ret := service.Update(req)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Delete(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service := api.GetService()
	err, ret := service.Delete(reqId.ID)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) DeleteByIds(c *gin.Context) {

	var reqIds request.GetByIds
	_ = c.ShouldBindJSON(&reqIds)
	if err := utils.Verify(reqIds, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	err, ret := service.DeleteByIds(reqIds.ID)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Find(c *gin.Context) {

	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	err, ret := service.Find(reqId.ID)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Detail(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	err, ret := service.Find(reqId.ID)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) ChangeStatus(c *gin.Context) {
	var req request.ChangeStatus
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service := api.GetService()
	err, ret := service.ChangeStatus(req.ID, req.Status)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Enable(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	service := api.GetService()
	var ids []uint64
	ids = append(ids, reqId.ID)
	err, ret := service.ChangeStatus(ids, true)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}

func (api *BaseApi) Disable(c *gin.Context) {

	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var ids []uint64
	ids = append(ids, reqId.ID)

	service := api.GetService()
	err, ret := service.ChangeStatus(ids, false)

	if err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}

	response.OkWithData(ret, c)
}
