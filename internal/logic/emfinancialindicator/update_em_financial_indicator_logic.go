package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmFinancialIndicatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmFinancialIndicatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmFinancialIndicatorLogic {
	return &UpdateEmFinancialIndicatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmFinancialIndicatorLogic) UpdateEmFinancialIndicator(req *types.EmFinancialIndicatorInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.UpdateEmFinancialIndicator(l.ctx,
		&efinance.EmFinancialIndicatorInfo{
			Id:                       req.Id,
			FinancialIndicatorsKey:   req.FinancialIndicatorsKey,
			FinancialIndicatorsValue: req.FinancialIndicatorsValue,
			FinancialIndicatorsName:  req.FinancialIndicatorsName,
			CreateTime:               req.CreateTime,
			Remarks:                  req.Remarks,
			Status:                   req.Status,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
