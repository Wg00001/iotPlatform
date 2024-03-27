package logic

import (
	"context"
	"iotPlatform/device/rpc/types/device"
	"iotPlatform/open/api/internal/svc"
	"iotPlatform/open/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageRequest) (resp string, err error) {
	_, err = l.svcCtx.RpcDevice.SendMessage(l.ctx, &device.SendMessageRequest{
		ProductKey: req.ProductKey,
		DeviceKey:  req.DeviceKey,
		Data:       req.Data,
	})
	if err != nil {
		logx.Error("ERR: open.logic.SendMessage: ", err)
	}
	return
}
