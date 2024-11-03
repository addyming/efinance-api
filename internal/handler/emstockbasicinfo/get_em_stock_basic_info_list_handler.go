package emstockbasicinfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockbasicinfo"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_basic_info/list emstockbasicinfo GetEmStockBasicInfoList
//
// Get em stock basic info list | 获取EmStockBasicInfo信息列表
//
// Get em stock basic info list | 获取EmStockBasicInfo信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmStockBasicInfoListReq
//
// Responses:
//  200: EmStockBasicInfoListResp

func GetEmStockBasicInfoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmStockBasicInfoListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockbasicinfo.NewGetEmStockBasicInfoListLogic(r.Context(), svcCtx)
		resp, err := l.GetEmStockBasicInfoList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
