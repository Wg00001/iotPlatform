package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"iotPlatform/common"
	"iotPlatform/common/jwts"
	"iotPlatform/models"
	"iotPlatform/user/api/internal/svc"
	"iotPlatform/user/api/internal/types"
	"log"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (token string, err error) {
	ub := new(models.UserBasic)
	err = l.svcCtx.DB.Where("name = ? AND password = ?", req.UserName, req.Password).First(&ub).Error
	if err != nil {
		log.Println("ERR: api.logic.user_login_logic.UserLogin: ", err)
		return "", errors.New("用户不存在")
	}
	token, err = jwts.GenToken(jwts.JwtPayLoad{
		Identity: ub.Id,
		Name:     ub.Name,
		Password: ub.Password,
	}, common.JwtAccessSecret, common.JwtAccessExpire)
	if err != nil {
		log.Println("ERR: api.logic.user_login_logic.UserLogin: ", err)
		return "", errors.New("token创建失败")
	}
	return token, err
}
