package emfinancialindicator

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_financial_indicator/update emfinancialindicator UpdateEmFinancialIndicator
//
// Update em financial indicator information | 更新EmFinancialIndicator信息
//
// Update em financial indicator information | 更新EmFinancialIndicator信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmFinancialIndicatorInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateEmFinancialIndicatorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmFinancialIndicatorInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emfinancialindicator.NewUpdateEmFinancialIndicatorLogic(r.Context(), svcCtx)
		resp, err := l.UpdateEmFinancialIndicator(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
