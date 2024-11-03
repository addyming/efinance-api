package emfinancialindicator

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_financial_indicator/delete emfinancialindicator DeleteEmFinancialIndicator
//
// Delete em financial indicator information | 删除EmFinancialIndicator信息
//
// Delete em financial indicator information | 删除EmFinancialIndicator信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsInt32Req
//
// Responses:
//  200: BaseMsgResp

func DeleteEmFinancialIndicatorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsInt32Req
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emfinancialindicator.NewDeleteEmFinancialIndicatorLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEmFinancialIndicator(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
