package emstockbasicinfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockbasicinfo"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_basic_info emstockbasicinfo GetEmStockBasicInfoById
//
// Get em stock basic info by ID | 通过ID获取EmStockBasicInfo信息
//
// Get em stock basic info by ID | 通过ID获取EmStockBasicInfo信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDInt32Req
//
// Responses:
//  200: EmStockBasicInfoInfoResp

func GetEmStockBasicInfoByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDInt32Req
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockbasicinfo.NewGetEmStockBasicInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetEmStockBasicInfoById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
