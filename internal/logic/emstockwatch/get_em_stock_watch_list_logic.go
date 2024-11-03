package emstockwatch

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockWatchListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmStockWatchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockWatchListLogic {
	return &GetEmStockWatchListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmStockWatchListLogic) GetEmStockWatchList(req *types.EmStockWatchListReq) (resp *types.EmStockWatchListResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.GetEmStockWatchList(l.ctx,
		&efinance.EmStockWatchListReq{
			Page:              req.Page,
			PageSize:          req.PageSize,
			StockWatchCode:    req.StockWatchCode,
			StockWatchName:    req.StockWatchName,
			StockWatchUpPrice: req.StockWatchUpPrice,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.EmStockWatchListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.EmStockWatchInfo{
				Id:                   v.Id,
				StockWatchCode:       v.StockWatchCode,
				StockWatchName:       v.StockWatchName,
				StockWatchUpPrice:    v.StockWatchUpPrice,
				StockWatchFallPrice:  v.StockWatchFallPrice,
				StockWatchUpChange:   v.StockWatchUpChange,
				StockWatchFallChange: v.StockWatchFallChange,
			})
	}
	return resp, nil
}
