package emstockwatch

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockwatch"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_watch/create emstockwatch CreateEmStockWatch
//
// Create em stock watch information | 创建EmStockWatch信息
//
// Create em stock watch information | 创建EmStockWatch信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmStockWatchInfo
//
// Responses:
//  200: BaseMsgResp

func CreateEmStockWatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmStockWatchInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockwatch.NewCreateEmStockWatchLogic(r.Context(), svcCtx)
		resp, err := l.CreateEmStockWatch(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
