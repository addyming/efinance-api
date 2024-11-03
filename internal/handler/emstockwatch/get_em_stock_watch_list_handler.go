package emstockwatch

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockwatch"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_watch/list emstockwatch GetEmStockWatchList
//
// Get em stock watch list | 获取EmStockWatch信息列表
//
// Get em stock watch list | 获取EmStockWatch信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmStockWatchListReq
//
// Responses:
//  200: EmStockWatchListResp

func GetEmStockWatchListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmStockWatchListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockwatch.NewGetEmStockWatchListLogic(r.Context(), svcCtx)
		resp, err := l.GetEmStockWatchList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
