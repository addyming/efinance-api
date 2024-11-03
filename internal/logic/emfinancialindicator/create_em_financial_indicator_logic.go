package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmFinancialIndicatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmFinancialIndicatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmFinancialIndicatorLogic {
	return &CreateEmFinancialIndicatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmFinancialIndicatorLogic) CreateEmFinancialIndicator(req *types.EmFinancialIndicatorInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.CreateEmFinancialIndicator(l.ctx,
		&efinance.EmFinancialIndicatorInfo{
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
