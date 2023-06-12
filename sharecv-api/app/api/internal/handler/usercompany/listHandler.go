package usercompany

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/usercompany"
	"sharecvapi/app/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := usercompany.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()

		response.Response(r, w, resp, err)
	}
}
