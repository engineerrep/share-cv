package config

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/config"
	"sharecvapi/app/api/internal/svc"
)

func GetConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := config.NewGetConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetConfig()

		response.Response(r, w, resp, err)
	}
}
