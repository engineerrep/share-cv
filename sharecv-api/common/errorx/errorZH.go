package errorx

var errorMsgZH map[int]string

func init() {
	configMsgZH()
}

func configMsgZH() {
	errorMsgZH = make(map[int]string)

	errorMsgZH[ServerErrorCode] = "服务繁忙，请稍后重试"
	errorMsgZH[CodeErrorCode] = "验证码错误"
	errorMsgEN[AccessErrorCode] = "无访问权限"
	errorMsgZH[EmailErrorCode] = "邮箱错误"
	errorMsgZH[PasswordErrorCode] = "密码错误"
	errorMsgZH[ParamErrorCode] = "参数错误"
	errorMsgZH[LoginErrorCode] = "登录失败"
	errorMsgZH[UploadErrorCode] = "上传失败"
	errorMsgZH[ChangeErrorCode] = "更改失败"
	errorMsgZH[SaveErrorCode] = "保存失败"
	errorMsgZH[EmailNotExistCode] = "邮箱不存在"
	errorMsgZH[TokenErrorCode] = "Token不存在"
	errorMsgZH[LogoutErrorCode] = "注销失败"
	errorMsgZH[InvitationCodeErrorCode] = "邀请码错误"
	errorMsgZH[ExistErrorCode] = "已存在"
	errorMsgZH[NotFountErrorCode] = "未找到"
	errorMsgZH[EmailExistCode] = "邮箱已存在"
	errorMsgZH[CodeSendErrorCode] = "验证码发送错误"
	errorMsgZH[RegisterErrorCode] = "注册失败"
	errorMsgZH[MaxCountErrorCode] = "超出最大限制"
}
