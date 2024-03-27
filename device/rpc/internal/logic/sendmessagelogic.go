package logic

import (
	"context"
	"errors"
	"iotPlatform/device/rpc/internal/mqtt"

	"iotPlatform/device/rpc/internal/svc"
	"iotPlatform/device/rpc/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	if in.ProductKey == "" || in.DeviceKey == "" || in.Data == "" {
		return nil, errors.New("参数异常")
	}
	topic := "/sys/" + in.ProductKey + "/" + in.DeviceKey + "/receive"
	if token := mqtt.MC.Publish(topic, 0, false, in.Data); token.Wait() && token.Error() != nil {
		logx.Error("ERR: device.rpc.logic.SendMessage : ", token.Error())
	}
	return &device.SendMessageReply{}, nil
}
