package login

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/token"
	"sharecvapi/common/utils"
	"sharecvapi/common/verify"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailPasswordLogic {
	return &EmailPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailPasswordLogic) EmailPassword(req *types.LoginEmailPasswordReq) (resp *types.LoginResp, err error) {
	// 邮箱是否注册
	user, err := l.svcCtx.AppUserModel.FindByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.EmailNotExistCode)
	}

	// 密码处理
	salt := fmt.Sprintf("%s%s", req.Password, l.svcCtx.Config.Salt)
	pwd := utils.MD5(salt)
	if user.Password != pwd {
		return nil, errorx.NewCodeError(errorx.PasswordErrorCode)
	}

	// 谷歌人机验证
	err = verify.GoogleVerifyReCaptchaV3(req.GoogleCaptcha, l.svcCtx.Config.Google.Recaptcha)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
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
	loginLog, err := l.svcCtx.AppUserLoginLogModel.FindByPlatform(l.ctx, user.Id, req.Platform)
	if err != nil {
		loginLog = &model.AppUserLoginLog{UserId: user.Id}
	}
	loginLog.Ip = ctxdata.GetUserIPFromCtx(l.ctx)
	loginLog.DeviceId = req.DeviceId
	loginLog.AppVersion = req.AppVersion
	loginLog.SystemVersion = req.SystemVersion
	loginLog.Platform = req.Platform

	l.svcCtx.AppUserLoginLogModel.Insert(l.ctx, loginLog)

	var val types.User
	err = copier.Copy(&val, &user)
	if err != nil {
		return nil, errorx.NewSystemError()
	}

	val.UserId = user.Uuid

	return &types.LoginResp{
		UserId:       val.UserId,
		User:         val,
		AccessToken:  token,
		AccessExpire: iat + l.svcCtx.Config.JwtAuth.AccessExpire,
	}, nil
}
