package emfinancialindicator

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_financial_indicator/list emfinancialindicator GetEmFinancialIndicatorList
//
// Get em financial indicator list | 获取EmFinancialIndicator信息列表
//
// Get em financial indicator list | 获取EmFinancialIndicator信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmFinancialIndicatorListReq
//
// Responses:
//  200: EmFinancialIndicatorListResp

func GetEmFinancialIndicatorListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmFinancialIndicatorListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emfinancialindicator.NewGetEmFinancialIndicatorListLogic(r.Context(), svcCtx)
		resp, err := l.GetEmFinancialIndicatorList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
