package emstockwatch

import (
	"context"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockWatchByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmStockWatchByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockWatchByIdLogic {
	return &GetEmStockWatchByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmStockWatchByIdLogic) GetEmStockWatchById(req *types.IDInt32Req) (resp *types.EmStockWatchInfoResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.GetEmStockWatchById(l.ctx, &efinance.IDInt32Req{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.EmStockWatchInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.EmStockWatchInfo{
			Id:                   data.Id,
			StockWatchCode:       data.StockWatchCode,
			StockWatchName:       data.StockWatchName,
			StockWatchUpPrice:    data.StockWatchUpPrice,
			StockWatchFallPrice:  data.StockWatchFallPrice,
			StockWatchUpChange:   data.StockWatchUpChange,
			StockWatchFallChange: data.StockWatchFallChange,
		},
	}, nil
}
