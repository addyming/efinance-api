package emstockwatch

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmStockWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmStockWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmStockWatchLogic {
	return &UpdateEmStockWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmStockWatchLogic) UpdateEmStockWatch(req *types.EmStockWatchInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.UpdateEmStockWatch(l.ctx,
		&efinance.EmStockWatchInfo{
			Id:                   req.Id,
			StockWatchCode:       req.StockWatchCode,
			StockWatchName:       req.StockWatchName,
			StockWatchUpPrice:    req.StockWatchUpPrice,
			StockWatchFallPrice:  req.StockWatchFallPrice,
			StockWatchUpChange:   req.StockWatchUpChange,
			StockWatchFallChange: req.StockWatchFallChange,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
