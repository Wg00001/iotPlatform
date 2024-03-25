package logic

import (
	"context"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductModifyLogic {
	return &ProductModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductModifyLogic) ProductModify(req *types.ProductModifyRequest) (resp string, err error) {
	err = l.svcCtx.DB.Where("id = ?", req.Id).
		Updates(&models.ProductBasic{
			Name: req.Name,
			Desc: req.Desc,
		}).Error
	if err != nil {
		log.Println("ERR: admin.api.logic.product_modify_logic: ERR1: ", err)
	}
	return
}
