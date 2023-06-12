package config

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/mmdb"
	"sharecvapi/common/utils"
)

type GetConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigLogic {
	return &GetConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigLogic) GetConfig() (resp *types.ConfigResp, err error) {

	// 默认美国服务器
	var config = types.ConfigResp{
		ApiUrls:    l.svcCtx.Config.App.ApiUrls,
		WebUrls:    l.svcCtx.Config.App.WebUrls,
		OssAuthUrl: l.svcCtx.Config.AliOSS.AuthUrl,
	}

	if l.svcCtx.Config.Mode == "dev" || l.svcCtx.Config.Mode == "pre" {
		config.OssBucketName = l.svcCtx.Config.AliOSS.BucketNameHK
		config.OssBucketUrl = l.svcCtx.Config.AliOSS.BucketUrlHK
		config.OssBucketEndpoint = l.svcCtx.Config.AliOSS.BucketEndpointHK
		return &config, nil
	}

	var realIP = ctxdata.GetUserIPFromCtx(l.ctx)
	oceania, country, err := mmdb.SearchContinent(realIP)
	if err != nil {
		return &config, nil
	}

	l.Logger.Infof(fmt.Sprintf("real ip: %s, oceania: %s, country : %s", realIP, oceania, country))

	// 判断是否是中国IP
	if country == l.svcCtx.Config.Region.China {
		config.OssBucketName = l.svcCtx.Config.AliOSS.BucketNameAP
		config.OssBucketUrl = l.svcCtx.Config.AliOSS.BucketUrlAP
		config.OssBucketEndpoint = l.svcCtx.Config.AliOSS.BucketEndpointAP
		return &config, nil
	}

	// 亚洲国家使用新加坡服务器
	if utils.StringIn(oceania, l.svcCtx.Config.Region.Asia) {
		config.OssBucketName = l.svcCtx.Config.AliOSS.BucketNameAP
		config.OssBucketUrl = l.svcCtx.Config.AliOSS.BucketUrlAP
		config.OssBucketEndpoint = l.svcCtx.Config.AliOSS.BucketEndpointAP
		return &config, nil
	}

	config.OssBucketName = l.svcCtx.Config.AliOSS.BucketNameUS
	config.OssBucketUrl = l.svcCtx.Config.AliOSS.BucketUrlUS
	config.OssBucketEndpoint = l.svcCtx.Config.AliOSS.BucketEndpointUS

	return &config, nil
}
