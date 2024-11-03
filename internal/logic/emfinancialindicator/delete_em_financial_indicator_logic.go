package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmFinancialIndicatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmFinancialIndicatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmFinancialIndicatorLogic {
	return &DeleteEmFinancialIndicatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmFinancialIndicatorLogic) DeleteEmFinancialIndicator(req *types.IDsInt32Req) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.DeleteEmFinancialIndicator(l.ctx, &efinance.IDsInt32Req{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
