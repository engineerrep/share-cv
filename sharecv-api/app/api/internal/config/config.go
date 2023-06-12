package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ChainRpc zrpc.RpcClientConf
	Salt     string

	App struct {
		Name       string
		IPDataPath string
		ApiUrls    []string
		WebUrls    []string
	}
	Region struct {
		China   string   // 中国大陆
		Asia    []string // 亚洲国家
		America []string // 美洲国家
		Europe  []string // 欧洲国家
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Ali struct {
		AccessKeyID     string
		AccessKeySecret string
	}
	AliEmail struct {
		AccountName string
		AddressType int32
		Endpoint    string
	}
	AliOSS struct {
		RoleArn          string
		RoleSessionName  string
		AuthUrl          string
		BucketNameHK     string
		BucketUrlHK      string
		BucketEndpointHK string
		BucketNameUS     string
		BucketUrlUS      string
		BucketEndpointUS string
		BucketNameAP     string
		BucketUrlAP      string
		BucketEndpointAP string
	}
	User struct {
		MaxPhoto           int
		MaxVideo           int
		MaxEducation       int
		MaxCustomAttribute int
		MaxInvitationCode  int
	}
	Google struct {
		Recaptcha     string
		AuthClientId  string
		AuthClientKey string
		State         string
		Scopes        []string
		AuthURL       string
		TokenURL      string
		UserInfoURL   string
	}
	OneSignal struct {
		Url        string
		AppId      string
		RestApiKey string
	}
	Twilio struct {
		AccountSid   string
		AccountToken string
		SmsFrom      string
	}
	SmtpEmail struct {
		FromAddress string
		FromName    string
		Subject     string
		Host        string
		Port        int
		Password    string
	}
	Mysql struct {
		DataSource string
	}
	Cache cache.CacheConf
	Redis struct {
		Host string
		Pass string
		Type string
	}
}
