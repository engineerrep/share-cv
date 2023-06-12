package userbase

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/userbase"
	"sharecvapi/app/api/internal/svc"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userbase.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()

		response.Response(r, w, resp, err)
	}
}
