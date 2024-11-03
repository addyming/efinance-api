package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmFinancialIndicatorByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmFinancialIndicatorByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmFinancialIndicatorByIdLogic {
	return &GetEmFinancialIndicatorByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmFinancialIndicatorByIdLogic) GetEmFinancialIndicatorById(req *types.IDInt32Req) (resp *types.EmFinancialIndicatorInfoResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.GetEmFinancialIndicatorById(l.ctx, &efinance.IDInt32Req{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.EmFinancialIndicatorInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.EmFinancialIndicatorInfo{
			Id:                       data.Id,
			FinancialIndicatorsKey:   data.FinancialIndicatorsKey,
			FinancialIndicatorsValue: data.FinancialIndicatorsValue,
			FinancialIndicatorsName:  data.FinancialIndicatorsName,
			CreateTime:               data.CreateTime,
			Remarks:                  data.Remarks,
			Status:                   data.Status,
		},
	}, nil
}
