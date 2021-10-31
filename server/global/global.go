package global

import (
	"gin-myboot/utils/timer"
	"github.com/gin-gonic/gin"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"gin-myboot/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
	Redis  *redis.Client
	Config config.Server
	Viper  *viper.Viper
	//Logger    *oplogging.Logger
	Logger             *zap.Logger
	Timer              timer.Timer = timer.NewTimerTask()
	ConcurrencyControl             = &singleflight.Group{}
	Engine             *gin.Engine
)
