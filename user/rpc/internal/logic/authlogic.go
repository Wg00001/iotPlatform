package logic

import (
	"context"
	"errors"
	"iotPlatform/common/jwts"
	"iotPlatform/user/rpc/internal/svc"
	"iotPlatform/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	if in.Token == "" {
		return nil, errors.New("必填项不能为空")
	}
	userClaim, err := jwts.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthReply)
	resp.Id = userClaim.Id
	resp.Extend = map[string]string{
		"name": userClaim.Name,
	}
	return resp, nil
}
