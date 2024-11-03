package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockBasicInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmStockBasicInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockBasicInfoByIdLogic {
	return &GetEmStockBasicInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmStockBasicInfoByIdLogic) GetEmStockBasicInfoById(req *types.IDInt32Req) (resp *types.EmStockBasicInfoInfoResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.GetEmStockBasicInfoById(l.ctx, &efinance.IDInt32Req{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.EmStockBasicInfoInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.EmStockBasicInfoInfo{
			Id:         data.Id,
			F1:         data.F1,
			F2:         data.F2,
			F3:         data.F3,
			F4:         data.F4,
			F5:         data.F5,
			F6:         data.F6,
			F7:         data.F7,
			F8:         data.F8,
			F9:         data.F9,
			F10:        data.F10,
			F11:        data.F11,
			F12:        data.F12,
			F13:        data.F13,
			F14:        data.F14,
			F15:        data.F15,
			F16:        data.F16,
			F17:        data.F17,
			F18:        data.F18,
			F19:        data.F19,
			F20:        data.F20,
			F21:        data.F21,
			F22:        data.F22,
			F23:        data.F23,
			F24:        data.F24,
			F25:        data.F25,
			F26:        data.F26,
			F27:        data.F27,
			F28:        data.F28,
			F29:        data.F29,
			F30:        data.F30,
			CreateTime: data.CreateTime,
			Remarks:    data.Remarks,
			Status:     data.Status,
			St1:        data.St1,
			St2:        data.St2,
			St3:        data.St3,
			St4:        data.St4,
			St5:        data.St5,
		},
	}, nil
}
