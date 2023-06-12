package web

import (
	"context"
	"net/http"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoogleAuthUrlLogic struct {
	logx.Logger
	res    http.ResponseWriter
	req    *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoogleAuthUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext, res http.ResponseWriter, req *http.Request) *GoogleAuthUrlLogic {
	return &GoogleAuthUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		res:    res,
		req:    req,
	}
}

func (l *GoogleAuthUrlLogic) GoogleAuthUrl(req *types.WebGoogleAuthReq) (resp *types.WebGoogleAuthResp, err error) {
	//oauth2Conf := &oauth2.Config{
	//	ClientID:     l.svcCtx.Config.Google.AuthClientId,
	//	ClientSecret: l.svcCtx.Config.Google.AuthClientKey,
	//	Scopes:       l.svcCtx.Config.Google.Scopes,
	//	Endpoint: oauth2.Endpoint{
	//		AuthURL:   l.svcCtx.Config.Google.AuthURL,
	//		TokenURL:  l.svcCtx.Config.Google.TokenURL,
	//		AuthStyle: oauth2.AuthStyleInParams,
	//	},
	//	RedirectURL: req.CallbackURL,
	//}

	//googleAuthUrl := oauth2Conf.AuthCodeURL(l.svcCtx.Config.Google.State)

	return &types.WebGoogleAuthResp{
		GoogleUrl:   l.svcCtx.Config.Google.AuthURL,
		ClientId:    l.svcCtx.Config.Google.AuthClientId,
		RedirectUri: req.CallbackURL,
		Scope:       strings.Join(l.svcCtx.Config.Google.Scopes, "+"),
		State:       l.svcCtx.Config.Google.State,
	}, nil
}
