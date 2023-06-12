package errorx

import "fmt"

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) ErrCode() int {
	return e.Code
}
func (e *CodeError) ErrorMsg() string {
	return e.Msg
}
func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode: %d", e.Code)
}

func newCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewCodeError(code int) error {
	return newCodeError(code, "")
}

func NewSystemError() error {
	return newCodeError(ServerErrorCode, "")
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
