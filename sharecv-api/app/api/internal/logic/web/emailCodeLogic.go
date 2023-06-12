package web

import (
	"context"
	"fmt"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"
	"sharecvapi/common/token"
	"sharecvapi/common/utils"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeLogic {
	return &EmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeLogic) EmailCode(req *types.WebLoginEmailCodeReq) (resp *types.WebLoginResp, err error) {
	// 邮箱验证码
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, req.Email)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	// 邮箱是否注册
	user, err := l.svcCtx.AppUserModel.FindByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.EmailNotExistCode)
	}

	// JWT登录鉴权
	iat := utils.TimeNow().Unix()
	token, err := token.GenToken(user.Id, user.Uuid, iat)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	// 判断自己是否存在邀请码
	if len(user.InvitationCode) != l.svcCtx.Config.User.MaxInvitationCode {
		user.InvitationCode = utils.InvitationCode(user.Id, l.svcCtx.Config.User.MaxInvitationCode)
		l.svcCtx.AppUserModel.Update(l.ctx, user)
	}

	// 登录日志
	loginLog, err := l.svcCtx.AppUserLoginLogModel.FindByPlatform(l.ctx, user.Id, 1)
	if err != nil {
		loginLog = &model.AppUserLoginLog{UserId: user.Id}
	}
	loginLog.Ip = ctxdata.GetUserIPFromCtx(l.ctx)
	loginLog.DeviceId = req.DeviceId
	loginLog.AppVersion = "web 1.0.0"
	loginLog.SystemVersion = req.BrowserName
	loginLog.Platform = 1

	l.svcCtx.AppUserLoginLogModel.Insert(l.ctx, loginLog)

	// 删除验证码缓存
	l.svcCtx.Redis.DelCtx(l.ctx, cacheCodeKey)

	var avatar string
	var firstName string
	var lastName string
	var nickname string

	userInfo, _ := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, user.Id)
	if userInfo != nil {
		avatar = userInfo.Avatar
		firstName = userInfo.FirstName
		lastName = userInfo.LastName
		nickname = userInfo.Nickname
	}

	return &types.WebLoginResp{
		UserId:       user.Uuid,
		AccessToken:  token,
		Email:        user.Email,
		Phone:        user.Mobile,
		PhoneNum:     user.MobileCountryNum,
		CountryId:    user.MobileCountryId,
		Nickname:     nickname,
		FirstName:    firstName,
		LastName:     lastName,
		Avatar:       avatar,
		AccessExpire: iat + l.svcCtx.Config.JwtAuth.AccessExpire,
	}, nil
}
