package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"iotPlatform/admin/api/internal/config"
	"iotPlatform/models"
	"iotPlatform/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RpcUser userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      models.NewDB(),
		RpcUser: userclient.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
