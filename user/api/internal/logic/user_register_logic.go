package logic

import (
	"context"
	"errors"
	"iotPlatform/common"
	"iotPlatform/common/jwts"
	"iotPlatform/models"
	"log"

	"iotPlatform/user/api/internal/svc"
	"iotPlatform/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserLoginRequest) (token string, err error) {
	ub := models.UserBasic{
		Name:     req.UserName,
		Password: req.Password,
	}
	res := l.svcCtx.DB.Create(&ub)
	if res.Error != nil {
		log.Println("ERR: api.logic.user_login_logic.UserLogin: ", err)
		return "", errors.New("用户已存在")
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
