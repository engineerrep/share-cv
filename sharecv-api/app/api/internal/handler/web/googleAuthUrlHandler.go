package web

import (
	"net/http"
	"sharecvapi/common/errorx"
	"sharecvapi/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sharecvapi/app/api/internal/logic/web"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
)

func GoogleAuthUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebGoogleAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, errorx.NewCodeError(errorx.ParamErrorCode))
			return
		}

		l := web.NewGoogleAuthUrlLogic(r.Context(), svcCtx, w, r)
		resp, err := l.GoogleAuthUrl(&req)

		response.Response(r, w, resp, err)
	}
}
