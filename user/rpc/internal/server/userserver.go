// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"iotPlatform/user/rpc/internal/logic"
	"iotPlatform/user/rpc/internal/svc"
	"iotPlatform/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Auth(ctx context.Context, in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	l := logic.NewAuthLogic(ctx, s.svcCtx)
	return l.Auth(in)
}

func (s *UserServer) OpenAuth(ctx context.Context, in *user.OpenAuthRequest) (*user.OpenAuthReply, error) {
	l := logic.NewOpenAuthLogic(ctx, s.svcCtx)
	return l.OpenAuth(in)
}