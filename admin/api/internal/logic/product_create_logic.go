package logic

import (
	"context"
	"github.com/google/uuid"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCreateLogic) ProductCreate(req *types.ProductCreateRequest) (resp string, err error) {
	err = l.svcCtx.DB.Create(&models.ProductBasic{
		Key:  uuid.New().String(),
		Name: req.Name,
		Desc: req.Desc,
	}).Error
	if err != nil {
		log.Println("ERR: admin.api.logic.product_create_logic: ERR1: ", err)
	}
	return
}
