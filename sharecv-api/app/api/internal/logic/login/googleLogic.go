package login

import (
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/token"
	"sharecvapi/common/utils"
	"sharecvapi/common/verify"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoogleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoogleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoogleLogic {
	return &GoogleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoogleLogic) Google(req *types.LoginGoogleReq) (resp *types.LoginResp, err error) {
	err, googleRep := verify.GoogleVerifyToken(l.ctx, req.IdToken, req.AccessToken)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	user, err := l.svcCtx.AppUserModel.FindByGoogleId(l.ctx, googleRep.UserId)
	// 邮箱未注册
	if user == nil {
		uid := utils.MD5(uuid.NewString())
		user = &model.AppUser{
			Email:     googleRep.Email,
			Uuid:      uid,
			GoogleUid: googleRep.UserId,
			UserType:  3,
		}

		// 注册一个新用户
		ret, err := l.svcCtx.AppUserModel.Insert(l.ctx, user)
		if err != nil {
			return nil, errorx.NewCodeError(errorx.LoginErrorCode)
		}

		user.CreateAt = utils.TimeNow()
		user.Id, err = ret.LastInsertId()
		if err != nil {
			return nil, errorx.NewCodeError(errorx.LoginErrorCode)
		}

		// 生成邀请码
		user.InvitationCode = utils.InvitationCode(user.Id, l.svcCtx.Config.User.MaxInvitationCode)
		l.svcCtx.AppUserModel.Update(l.ctx, user)

		// 随机生成一个昵称
		userInfo := &model.AppUserInfo{
			UserId:   user.Id,
			Nickname: utils.RandString(10),
		}
		l.svcCtx.AppUserInfoModel.Insert(l.ctx, userInfo)
	}

	// JWT登录鉴权
	iat := utils.TimeNow().Unix()
	token, err := token.GenToken(user.Id, user.Uuid, iat)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	// 邀请码不存在
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
