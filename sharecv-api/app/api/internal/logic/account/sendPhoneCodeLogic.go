package account

import (
	"context"
	"fmt"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"
	"sharecvapi/common/notice"
	"sharecvapi/common/utils"
	"sharecvapi/common/verify"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendPhoneCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneCodeLogic {
	return &SendPhoneCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendPhoneCodeLogic) SendPhoneCode(req *types.SendCodePhoneReq) (resp *types.CodeResp, err error) {
	// 谷歌人机验证
	err = verify.GoogleVerifyReCaptchaV3(req.GoogleCaptcha, l.svcCtx.Config.Google.Recaptcha)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	phone := fmt.Sprintf("+%s%s", req.PhoneNum, req.Phone)
	var phones []string
	phones = append(phones, phone)
	// 获取缓存Code
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, phone)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	language := ctxdata.GetLanguageFromCtx(l.ctx)
	var retCode int
	if len(verifyCode) == 6 {
		// 发送邮箱验证码
		retCode, _ = notice.OnesignalSMS(language, l.svcCtx.Config.App.Name, phones, l.svcCtx.Config.Twilio.SmsFrom, verifyCode) //.SendTwilioSMS(language, verifyCode, phone)
	} else {
		// 6位随机验证码
		code := utils.RandomInt(100000, 999999)
		verifyCode = fmt.Sprintf("%d", code)

		err = l.svcCtx.Redis.SetexCtx(l.ctx, cacheCodeKey, verifyCode, globalkey.CacheUserCodeTtl)
		if err != nil {
			return nil, errorx.NewCodeError(errorx.CodeSendErrorCode)
		}

		// 发送邮箱验证码
		retCode, _ = notice.OnesignalSMS(language, l.svcCtx.Config.App.Name, phones, l.svcCtx.Config.Twilio.SmsFrom, verifyCode) //.SendTwilioSMS(language, verifyCode, phone)
	}

	if retCode != 200 {
		l.svcCtx.Redis.DelCtx(l.ctx, cacheCodeKey)
		return nil, errorx.NewCodeError(errorx.CodeSendErrorCode)
	}
	code := ""
	if l.svcCtx.Config.Mode == "dev" {
		code = verifyCode
	}
	return &types.CodeResp{
		Status: 1,
		Code:   code,
	}, nil
}
