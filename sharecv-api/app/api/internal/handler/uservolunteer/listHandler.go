package uservolunteer

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/uservolunteer"
	"sharecvapi/app/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := uservolunteer.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()

		response.Response(r, w, resp, err)
	}
}
