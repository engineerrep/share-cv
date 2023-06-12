package oss

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/oss"
	"sharecvapi/app/api/internal/svc"
)

func OssAuthServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oss.NewOssAuthServerLogic(r.Context(), svcCtx)
		resp, err := l.OssAuthServer()

		response.Response(r, w, resp, err)
	}
}
