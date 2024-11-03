package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockBasicInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmStockBasicInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockBasicInfoListLogic {
	return &GetEmStockBasicInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmStockBasicInfoListLogic) GetEmStockBasicInfoList(req *types.EmStockBasicInfoListReq) (resp *types.EmStockBasicInfoListResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.GetEmStockBasicInfoList(l.ctx,
		&efinance.EmStockBasicInfoListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.EmStockBasicInfoListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.EmStockBasicInfoInfo{
				Id:         v.Id,
				F1:         v.F1,
				F2:         v.F2,
				F3:         v.F3,
				F4:         v.F4,
				F5:         v.F5,
				F6:         v.F6,
				F7:         v.F7,
				F8:         v.F8,
				F9:         v.F9,
				F10:        v.F10,
				F11:        v.F11,
				F12:        v.F12,
				F13:        v.F13,
				F14:        v.F14,
				F15:        v.F15,
				F16:        v.F16,
				F17:        v.F17,
				F18:        v.F18,
				F19:        v.F19,
				F20:        v.F20,
				F21:        v.F21,
				F22:        v.F22,
				F23:        v.F23,
				F24:        v.F24,
				F25:        v.F25,
				F26:        v.F26,
				F27:        v.F27,
				F28:        v.F28,
				F29:        v.F29,
				F30:        v.F30,
				CreateTime: v.CreateTime,
				Remarks:    v.Remarks,
				Status:     v.Status,
				St1:        v.St1,
				St2:        v.St2,
				St3:        v.St3,
				St4:        v.St4,
				St5:        v.St5,
			})
	}
	return resp, nil
}
