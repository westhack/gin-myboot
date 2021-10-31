package service

import (
	"context"
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/system/model"
	"gin-myboot/utils"
	"time"
)

type RedisService struct {
}

var RedisServiceApp = new(RedisService)

// GetList
// @function: GetList
// @description: 分页获取数据
// @param: info request.QueryParams
// @return: err error, list interface{}, total int64
func (redisService *RedisService) GetList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)

	ctx := context.Background()
	key := "*"
	if queryParams.Search != nil && len(queryParams.Search) > 0 {
		v := queryParams.Search[0]
		if v.Value != nil {
			str := utils.Strval(v.Value)
			if str != "" {
				key = "*" + str + "*"
			}
		}
	}
	global.Error("=======>", key)

	results, err := global.Redis.Keys(ctx, key).Result()
	if err != nil {
		return err, nil, 0
	}
	global.Error("=======>", results)
	total = int64(len(results))

	if total > 0 {
		if int64(offset + limit) < total {
			results = results[offset:limit]
		} else {
			results = results[offset:]
		}
	}

	var ret []model.RedisInfo
	for _, k := range results {
		value, _ := global.Redis.Get(ctx, k).Result()

		t, _ := global.Redis.TTL(ctx, k).Result()

		ret = append(ret, model.RedisInfo{
			Key:        k,
			Value:      value,
			ExpireTime: t.Milliseconds() / 1000,
		})
	}

	return err, ret, total
}

// Create
// @function: Create
// @description: 创建缓存
// @param: info model.RedisInfo
// @return: err error
func (redisService *RedisService) Create(info model.RedisInfo) (err error) {
	ctx := context.Background()
	_, err = global.Redis.Set(ctx, info.Key, info.Value, time.Duration(info.ExpireTime)*time.Second).Result()
	if err != nil {
		return err
	}

	return err
}

// Update
// @function: Update
// @description: 创建缓存
// @param: info model.RedisInfo
// @return: err error
func (redisService *RedisService) Update(info model.RedisInfo) (err error) {
	ctx := context.Background()
	_, err = global.Redis.Set(ctx, info.Key, info.Value, time.Duration(info.ExpireTime)*time.Second).Result()
	if err != nil {
		return err
	}

	return err
}

// FindByKey
// @function: FindByKey
// @description: 根据key获取
// @param: info model.RedisInfo
// @return: err error
func (redisService *RedisService) FindByKey(key string) (info model.RedisInfo, err error) {
	ctx := context.Background()
	res, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		return info, err
	}

	t, _ := global.Redis.TTL(ctx, key).Result()

	info.Key = key
	info.Value = res
	info.ExpireTime = t.Milliseconds() / 1000

	return info, err
}

// Delete
// @function: Delete
// @description: 创建缓存
// @param: info systemRequest.GetRedisKey
// @return: err error
func (redisService *RedisService) Delete(key string) (err error) {
	ctx := context.Background()
	_, err = global.Redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 创建缓存
// @param: info systemRequest.GetRedisKeys
// @return: err error
func (redisService *RedisService) DeleteByIds(keys []string) (err error) {
	ctx := context.Background()

	for _, key := range keys {
		_, err = global.Redis.Del(ctx, key).Result()
		if err != nil {
			return err
		}
	}

	return err
}
