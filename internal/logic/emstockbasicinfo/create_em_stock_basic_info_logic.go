package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmStockBasicInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmStockBasicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmStockBasicInfoLogic {
	return &CreateEmStockBasicInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmStockBasicInfoLogic) CreateEmStockBasicInfo(req *types.EmStockBasicInfoInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.CreateEmStockBasicInfo(l.ctx,
		&efinance.EmStockBasicInfoInfo{
			F1:         req.F1,
			F2:         req.F2,
			F3:         req.F3,
			F4:         req.F4,
			F5:         req.F5,
			F6:         req.F6,
			F7:         req.F7,
			F8:         req.F8,
			F9:         req.F9,
			F10:        req.F10,
			F11:        req.F11,
			F12:        req.F12,
			F13:        req.F13,
			F14:        req.F14,
			F15:        req.F15,
			F16:        req.F16,
			F17:        req.F17,
			F18:        req.F18,
			F19:        req.F19,
			F20:        req.F20,
			F21:        req.F21,
			F22:        req.F22,
			F23:        req.F23,
			F24:        req.F24,
			F25:        req.F25,
			F26:        req.F26,
			F27:        req.F27,
			F28:        req.F28,
			F29:        req.F29,
			F30:        req.F30,
			CreateTime: req.CreateTime,
			Remarks:    req.Remarks,
			Status:     req.Status,
			St1:        req.St1,
			St2:        req.St2,
			St3:        req.St3,
			St4:        req.St4,
			St5:        req.St5,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
