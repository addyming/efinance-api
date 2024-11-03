package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmFinancialIndicatorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmFinancialIndicatorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmFinancialIndicatorListLogic {
	return &GetEmFinancialIndicatorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmFinancialIndicatorListLogic) GetEmFinancialIndicatorList(req *types.EmFinancialIndicatorListReq) (resp *types.EmFinancialIndicatorListResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.GetEmFinancialIndicatorList(l.ctx,
		&efinance.EmFinancialIndicatorListReq{
			Page:                     req.Page,
			PageSize:                 req.PageSize,
			FinancialIndicatorsKey:   req.FinancialIndicatorsKey,
			FinancialIndicatorsValue: req.FinancialIndicatorsValue,
			FinancialIndicatorsName:  req.FinancialIndicatorsName,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.EmFinancialIndicatorListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.EmFinancialIndicatorInfo{
				Id:                       v.Id,
				FinancialIndicatorsKey:   v.FinancialIndicatorsKey,
				FinancialIndicatorsValue: v.FinancialIndicatorsValue,
				FinancialIndicatorsName:  v.FinancialIndicatorsName,
				CreateTime:               v.CreateTime,
				Remarks:                  v.Remarks,
				Status:                   v.Status,
			})
	}
	return resp, nil
}
