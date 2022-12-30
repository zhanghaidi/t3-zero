package svc

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"t3-zero/common/model"
	"t3-zero/front/internal/config"
	"t3-zero/front/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		DB:     model.Init(c.Mysql.DataSource),
		RDB:    model.InitRedis(c.Redis.Addr),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
