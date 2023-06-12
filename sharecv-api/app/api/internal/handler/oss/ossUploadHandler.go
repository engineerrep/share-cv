package oss

import (
	"net/http"
	"sharecvapi/app/api/internal/logic/oss"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/common/response"
)

func OssUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oss.NewOssUploadLogic(r.Context(), svcCtx)
		resp, err := l.OssUpload(r)

		response.Response(r, w, resp, err)
	}
}
