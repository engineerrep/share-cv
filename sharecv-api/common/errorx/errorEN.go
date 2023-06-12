package errorx

var errorMsgEN map[int]string

func init() {
	configMsgEN()
}

func configMsgEN() {
	errorMsgEN = make(map[int]string)
	errorMsgEN[ServerErrorCode] = "Service is busy"
	errorMsgEN[CodeErrorCode] = "Code verification error"
	errorMsgEN[AccessErrorCode] = "Access not turned on"
	errorMsgEN[EmailErrorCode] = "Email error"
	errorMsgEN[PasswordErrorCode] = "Password error"
	errorMsgEN[ParamErrorCode] = "Params error"
	errorMsgEN[LoginErrorCode] = "Login error"
	errorMsgEN[UploadErrorCode] = "Upload error"
	errorMsgEN[ChangeErrorCode] = "Change error"
	errorMsgEN[EmailNotExistCode] = "Email not exist"
	errorMsgEN[SaveErrorCode] = "Save error"
	errorMsgEN[TokenErrorCode] = "Token is error"
	errorMsgEN[LogoutErrorCode] = "Logout is error"
	errorMsgEN[InvitationCodeErrorCode] = "InvitationCode error"
	errorMsgEN[ExistErrorCode] = "Exist error"
	errorMsgEN[NotFountErrorCode] = "Not Found"
	errorMsgEN[EmailExistCode] = "Email exist"
	errorMsgEN[CodeSendErrorCode] = "Send code error"
	errorMsgEN[RegisterErrorCode] = "Register error"
	errorMsgEN[MaxCountErrorCode] = "Maximum limit error"
}
