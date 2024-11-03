package emstockwatch

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockwatch"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_watch/delete emstockwatch DeleteEmStockWatch
//
// Delete em stock watch information | 删除EmStockWatch信息
//
// Delete em stock watch information | 删除EmStockWatch信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsInt32Req
//
// Responses:
//  200: BaseMsgResp

func DeleteEmStockWatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsInt32Req
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockwatch.NewDeleteEmStockWatchLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEmStockWatch(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
