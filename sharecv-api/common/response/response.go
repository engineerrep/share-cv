package response

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ResBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	body := new(ResBody)
	if err != nil {
		language := ctxdata.GetLanguageFromCtx(r.Context())
		causeErr := errors.Cause(err)
		errCode := errorx.ServerErrorCode
		errMsg := err.Error()
		// 自定义错误类型
		if e, ok := causeErr.(*errorx.CodeError); ok {
			errCode = e.ErrCode()
			errMsg = errorx.MapErrMsg(errCode, language)
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR : %v, errCode: %d, errMsg: %s \n\n", err, errCode, errMsg)
		body.Code = errCode
		body.Msg = errMsg
	} else {
		body.Code = 200
		body.Msg = "success"
		body.Data = resp
	}
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "*")
	httpx.OkJson(w, body)
}

// TokenErrorResponse token错误返回
func TokenErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	language := ctxdata.GetLanguageFromCtx(r.Context())

	body := new(ResBody)
	body.Code = errorx.TokenErrorCode
	body.Msg = errorx.MapErrMsg(errorx.TokenErrorCode, language)

	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "*")
	httpx.OkJson(w, body)
}
