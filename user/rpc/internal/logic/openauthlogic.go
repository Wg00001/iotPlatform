package logic

import (
	"context"

	"iotPlatform/user/rpc/internal/svc"
	"iotPlatform/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenAuthLogic {
	return &OpenAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenAuthLogic) OpenAuth(in *user.OpenAuthRequest) (*user.OpenAuthReply, error) {
	// todo: add your logic here and delete this line

	return &user.OpenAuthReply{}, nil
}
