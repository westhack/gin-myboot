package initialize

import (
	"context"
	"gin-myboot/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.Logger.Info("redis connect ping response:", zap.String("pong", pong))
		var ctx = context.Background()
		client.Set(ctx, "text", "!", 0)
		global.Redis = client
	}
}
