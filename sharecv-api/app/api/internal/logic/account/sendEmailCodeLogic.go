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

type SendEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(req *types.SendCodeEmailReq) (resp *types.CodeResp, err error) {

	// 谷歌人机验证
	err = verify.GoogleVerifyReCaptchaV3(req.GoogleCaptcha, l.svcCtx.Config.Google.Recaptcha)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	aliConf := notice.AliEmailConfig{AccessKeyID: l.svcCtx.Config.Ali.AccessKeyID,
		AccessKeySecret: l.svcCtx.Config.Ali.AccessKeySecret,
		AccountName:     l.svcCtx.Config.AliEmail.AccountName,
		AddressType:     l.svcCtx.Config.AliEmail.AddressType,
		Endpoint:        l.svcCtx.Config.AliEmail.Endpoint,
		AppName:         l.svcCtx.Config.App.Name,
	}
	// 获取缓存Code
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, req.Email)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	language := ctxdata.GetLanguageFromCtx(l.ctx)
	var retCode int32
	if len(verifyCode) == 6 {
		// 发送邮箱验证码
		retCode, _ = notice.AliEmailCode(aliConf, language, verifyCode, req.Email)
	} else {
		// 6位随机验证码
		code := utils.RandomInt(100000, 999999)
		verifyCode = fmt.Sprintf("%d", code)

		err = l.svcCtx.Redis.SetexCtx(l.ctx, cacheCodeKey, verifyCode, globalkey.CacheUserCodeTtl)
		if err != nil {
			return nil, errorx.NewCodeError(errorx.CodeSendErrorCode)
		}

		// 发送邮箱验证码
		retCode, _ = notice.AliEmailCode(aliConf, language, verifyCode, req.Email)
	}

	if retCode != 200 {
		l.svcCtx.Redis.DelCtx(l.ctx, cacheCodeKey)
		return nil, errorx.NewCodeError(errorx.CodeSendErrorCode)
	}

	return &types.CodeResp{
		Status: 1,
	}, nil
}
