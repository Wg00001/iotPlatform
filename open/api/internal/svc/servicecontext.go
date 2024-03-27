package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"iotPlatform/device/rpc/deviceclient"
	"iotPlatform/open/api/internal/config"
	"iotPlatform/user/rpc/userclient"
)

type ServiceContext struct {
	Config    config.Config
	RpcUser   userclient.User //引入rpc
	RpcDevice deviceclient.Device
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: deviceclient.NewDevice(zrpc.MustNewClient(c.RpcDevice)),
		RpcUser:   userclient.NewUser(zrpc.MustNewClient(c.RpcUser)),
	}
}
