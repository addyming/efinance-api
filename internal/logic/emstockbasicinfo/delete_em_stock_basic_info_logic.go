package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
	//"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmStockBasicInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmStockBasicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmStockBasicInfoLogic {
	return &DeleteEmStockBasicInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmStockBasicInfoLogic) DeleteEmStockBasicInfo(req *types.IDsInt32Req) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.EfinanceRpc.DeleteEmStockBasicInfo(l.ctx, &efinance.IDsInt32Req{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
