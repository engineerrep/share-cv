package errorx

import (
	"sharecvapi/common/ctxdata"
)

func MapErrMsg(errCode int, language string) string {
	switch language {
	case ctxdata.LanguageZH:
		if msg, ok := errorMsgZH[errCode]; ok {
			return msg
		}
		return "服务错误"
	default:
		if msg, ok := errorMsgEN[errCode]; ok {
			return msg
		}
		return "Service error"
	}
}
