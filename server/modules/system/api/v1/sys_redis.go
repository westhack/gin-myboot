package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/system/model"
	systemRequest "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedisApi struct {
}

// GetList
// @Tags SysRedis
// @Summary 分页获取缓存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchParams true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/redis/getList [post]
func (b *RedisApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := redisService.GetList(searchParams); err != nil {
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

// Create
// @Tags SysRedis
// @Summary 分页获取缓存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RedisInfo true "缓存数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/redis/create [post]
func (b *RedisApi) Create(c *gin.Context) {
	var redisInfo model.RedisInfo
	if err := c.ShouldBindJSON(&redisInfo); err != nil {
		global.Error("======>", redisInfo)
		response.FailWithMessage("创建失败，"+utils.GetError(err, redisInfo), c)
		return
	}

	if err := redisService.Create(redisInfo); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(gin.H{"redisInfo": redisInfo}, "创建成功", c)
	}
}

// Update
// @Tags SysRedis
// @Summary 分页获取缓存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RedisInfo true "缓存数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/redis/update [post]
func (b *RedisApi) Update(c *gin.Context) {
	var redisInfo model.RedisInfo
	if err := c.ShouldBindJSON(&redisInfo); err != nil {
		response.FailWithMessage("修改失败，"+utils.GetError(err, redisInfo), c)
		return
	}

	if err := redisService.Update(redisInfo); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithDetailed(gin.H{"redisInfo": redisInfo}, "修改成功", c)
	}
}

// Find
// @Tags SysRedis
// @Summary 分页获取缓存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "缓存key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/redis/find [get]
func (b *RedisApi) Find(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		response.FailWithMessage("key 不能为空", c)
		return
	}

	if res, err := redisService.FindByKey(key); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithDetailed(gin.H{"info": res}, "获取成功", c)
	}
}

// Delete
// @Tags SysRedis
// @Summary 分页获取缓存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemRequest.GetRedisKey true "缓存数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/redis/delete [post]
func (b *RedisApi) Delete(c *gin.Context) {
	var req systemRequest.GetRedisKey
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("删除失败，"+utils.GetError(err, req), c)
		return
	}

	if err := redisService.Delete(req.Key); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags SysRedis
// @Summary 分页获取缓存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemRequest.GetRedisKeys true "缓存数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/redis/deleteByIds [post]
func (b *RedisApi) DeleteByIds(c *gin.Context) {
	var req systemRequest.GetRedisKeys
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("删除失败，"+utils.GetError(err, req), c)
		return
	}

	if err := redisService.DeleteByIds(req.Key); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
