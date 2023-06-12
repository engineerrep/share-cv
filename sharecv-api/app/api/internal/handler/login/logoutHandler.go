package login

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/login"
	"sharecvapi/app/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()

		response.Response(r, w, resp, err)
	}
}
