package emstockbasicinfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockbasicinfo"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_basic_info/create emstockbasicinfo CreateEmStockBasicInfo
//
// Create em stock basic info information | 创建EmStockBasicInfo信息
//
// Create em stock basic info information | 创建EmStockBasicInfo信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmStockBasicInfoInfo
//
// Responses:
//  200: BaseMsgResp

func CreateEmStockBasicInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmStockBasicInfoInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockbasicinfo.NewCreateEmStockBasicInfoLogic(r.Context(), svcCtx)
		resp, err := l.CreateEmStockBasicInfo(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
