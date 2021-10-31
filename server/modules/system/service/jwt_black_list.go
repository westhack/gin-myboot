package service

import (
	"context"
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"time"

	"gorm.io/gorm"
)

type JwtService struct {
}

// JsonInBlacklist
// @author: [piexlmax](https://github.com/piexlmax)
// @function: JsonInBlacklist
// @description: 拉黑jwt
// @param: jwtList model.JwtBlacklist
// @return: err error
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GormDB.Create(&jwtList).Error
	return
}

// IsBlacklist
// @author: [piexlmax](https://github.com/piexlmax)
// @function: IsBlacklist
// @description: 判断JWT是否在黑名单内部
// @param: jwt string
// @return: bool
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.GormDB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

// GetRedisJWT
// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetRedisJWT
// @description: 从redis取jwt
// @param: userName string
// @return: err error, redisJWT string
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.Redis.Get(context.Background(), userName).Result()
	return err, redisJWT
}

// SetRedisJWT
// @author: [piexlmax](https://github.com/piexlmax)
// @function: SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param: jwt string, userName string
// @return: err error
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.Config.JWT.ExpiresTime) * time.Second
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
