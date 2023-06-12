package register

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"sharecvapi/app/model"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"
	"sharecvapi/common/utils"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterEmailLogic {
	return &RegisterEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterEmailLogic) RegisterEmail(req *types.RegisterEmailReq) (resp *types.StatusResp, err error) {
	// 邮箱验证码
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, req.Email)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	// 邮箱是否存在
	user, err := l.svcCtx.AppUserModel.FindByEmail(l.ctx, req.Email)
	if user != nil {
		return nil, errorx.NewCodeError(errorx.EmailExistCode)
	}

	uid := utils.MD5(uuid.NewString())
	salt := fmt.Sprintf("%s%s", req.Password, l.svcCtx.Config.Salt)
	pwd := utils.MD5(salt)
	user = &model.AppUser{
		Email:    req.Email,
		Uuid:     uid,
		Password: pwd,
		UserType: 1,
	}

	// 注册一个新用户
	ret, err := l.svcCtx.AppUserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.RegisterErrorCode)
	}

	user.CreateAt = utils.TimeNow()
	user.Id, err = ret.LastInsertId()
	if err != nil {
		return nil, errorx.NewCodeError(errorx.RegisterErrorCode)
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

	// 上级邀请码
	if len(req.InvitationCode) == l.svcCtx.Config.User.MaxInvitationCode {
		superUser, _ := l.svcCtx.AppUserModel.FindByInvitationCode(l.ctx, req.InvitationCode)
		if superUser != nil {
			userInvitation := &model.AppUserInvitation{UserId: user.Id, SuperUserId: superUser.Id, InvitationCode: req.InvitationCode}
			l.svcCtx.AppUserInvitationModel.Insert(l.ctx, userInvitation)
		}
	}

	// 删除验证码缓存
	l.svcCtx.Redis.DelCtx(l.ctx, cacheCodeKey)

	return &types.StatusResp{
		Status: 1,
	}, nil
}
