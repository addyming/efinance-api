package emfinancialindicator

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_financial_indicator/create emfinancialindicator CreateEmFinancialIndicator
//
// Create em financial indicator information | 创建EmFinancialIndicator信息
//
// Create em financial indicator information | 创建EmFinancialIndicator信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmFinancialIndicatorInfo
//
// Responses:
//  200: BaseMsgResp

func CreateEmFinancialIndicatorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmFinancialIndicatorInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emfinancialindicator.NewCreateEmFinancialIndicatorLogic(r.Context(), svcCtx)
		resp, err := l.CreateEmFinancialIndicator(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
