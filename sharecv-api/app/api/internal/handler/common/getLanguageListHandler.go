package common

import (
	"net/http"
	"sharecvapi/common/response"

	"sharecvapi/app/api/internal/logic/common"
	"sharecvapi/app/api/internal/svc"
)

func GetLanguageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := common.NewGetLanguageListLogic(r.Context(), svcCtx)
		resp, err := l.GetLanguageList()

		response.Response(r, w, resp, err)
	}
}
