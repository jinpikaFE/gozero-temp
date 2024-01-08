package logic

import (
	"context"

	"gzero-user/service/product/api/internal/svc"
	"gzero-user/service/product/api/internal/types"
	"gzero-user/service/product/rpc/pb/gzero-user/service/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
        Id: req.Id,
    })
    if err != nil {
        return nil, err
    }

    return &types.DetailResponse{
        Id:     res.Id,
        Name:   res.Name,
        Desc:   res.Desc,
        Stock:  res.Stock,
        Amount: res.Amount,
        Status: res.Status,
    }, nil
}
