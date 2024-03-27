package svc

import (
	"gorm.io/gorm"
	"iotPlatform/models"
	"iotPlatform/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.NewDB(),
	}
}
