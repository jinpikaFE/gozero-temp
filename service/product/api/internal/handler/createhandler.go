package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gzero-user/service/product/api/internal/logic"
	"gzero-user/service/product/api/internal/svc"
	"gzero-user/service/product/api/internal/types"
	xhttp "github.com/zeromicro/x/http"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
