package web

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/oauth2"
	"io"
	"sharecvapi/app/api/internal/logic/oss"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/token"
	"sharecvapi/common/utils"
)

type GoogleTokenUser struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

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

func (l *GoogleLogic) Google(req *types.WebLoginGoogleReq) (resp *types.WebLoginResp, err error) {

	oauth2Conf := &oauth2.Config{
		ClientID:     l.svcCtx.Config.Google.AuthClientId,
		ClientSecret: l.svcCtx.Config.Google.AuthClientKey,
		Endpoint: oauth2.Endpoint{
			TokenURL:  l.svcCtx.Config.Google.TokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: req.CallbackURL,
	}

	// 客户端如果给Code需要去谷歌获取Token
	//gToken, err := oauth2Conf.Exchange(l.ctx, req.Code)
	//if err != nil {
	//	return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	//}

	// 客户端直接给谷歌Token
	gToken := &oauth2.Token{
		AccessToken: req.Code,
	}

	client := oauth2Conf.Client(l.ctx, gToken)
	response, err := client.Get(l.svcCtx.Config.Google.UserInfoURL)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	if response.StatusCode != 200 {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	gBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	gUser := GoogleTokenUser{}
	json.Unmarshal(gBody, &gUser)
	if len(gUser.Id) == 0 {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	var avatar string
	var firstName string
	var lastName string
	var nickname string

	user, err := l.svcCtx.AppUserModel.FindByGoogleId(l.ctx, gUser.Id)
	// 邮箱未注册
	if user == nil {
		uid := utils.MD5(uuid.NewString())
		user = &model.AppUser{
			Email:     gUser.Email,
			Uuid:      uid,
			GoogleUid: gUser.Id,
			UserType:  3,
			Locale:    gUser.Locale,
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

		// 把谷歌头像上传到OSS
		var googlePicture string
		googlePicture, err = oss.UploadWithImageUrl(l.ctx, l.svcCtx, gUser.Picture)
		if err != nil {
			googlePicture = gUser.Picture
		}

		// 创建用户详情
		userInfo := &model.AppUserInfo{
			UserId:    user.Id,
			FirstName: gUser.GivenName,
			LastName:  gUser.FamilyName,
			Avatar:    googlePicture,
			Nickname:  gUser.Name,
		}
		l.svcCtx.AppUserInfoModel.Insert(l.ctx, userInfo)

		avatar = gUser.Picture
		firstName = gUser.GivenName
		lastName = gUser.FamilyName
		nickname = gUser.Name
	} else {
		userInfo, _ := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, user.Id)
		if userInfo != nil {
			avatar = userInfo.Avatar
			firstName = userInfo.FirstName
			lastName = userInfo.LastName
			nickname = userInfo.Nickname
		}
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
