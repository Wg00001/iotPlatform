// Code generated by goctl. DO NOT EDIT.
// Source: device.proto

package deviceclient

import (
	"context"

	"iotPlatform/device/rpc/types/device"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendMessageReply   = device.SendMessageReply
	SendMessageRequest = device.SendMessageRequest

	Device interface {
		SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageReply, error)
	}

	defaultDevice struct {
		cli zrpc.Client
	}
)

func NewDevice(cli zrpc.Client) Device {
	return &defaultDevice{
		cli: cli,
	}
}

func (m *defaultDevice) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageReply, error) {
	client := device.NewDeviceClient(m.cli.Conn())
	return client.SendMessage(ctx, in, opts...)
}
