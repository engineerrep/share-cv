package usereducation

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sharecvapi/app/api/internal/logic/usereducation"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserEducationReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, errorx.NewCodeError(errorx.ParamErrorCode))
			return
		}

		validate := validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		language := ctxdata.GetLanguageFromCtx(r.Context())
		trans, _ := ut.New(zh.New()).GetTranslator(language)
		validateErr := translations.RegisterDefaultTranslations(validate, trans)
		if validateErr = validate.StructCtx(r.Context(), req); validateErr != nil {
			for _, err := range validateErr.(validator.ValidationErrors) {
				if err != nil {
					response.Response(r, w, nil, errorx.NewCodeError(errorx.ParamErrorCode))
					return
				}
			}
		}

		l := usereducation.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(&req)

		response.Response(r, w, resp, err)
	}
}
