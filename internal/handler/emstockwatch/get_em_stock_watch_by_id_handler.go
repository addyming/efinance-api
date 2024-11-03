package emstockwatch

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockwatch"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_watch emstockwatch GetEmStockWatchById
//
// Get em stock watch by ID | 通过ID获取EmStockWatch信息
//
// Get em stock watch by ID | 通过ID获取EmStockWatch信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDInt32Req
//
// Responses:
//  200: EmStockWatchInfoResp

func GetEmStockWatchByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDInt32Req
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockwatch.NewGetEmStockWatchByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetEmStockWatchById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
