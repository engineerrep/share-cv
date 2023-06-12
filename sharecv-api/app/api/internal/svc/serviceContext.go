package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sharecvapi/app/api/internal/config"
	"sharecvapi/app/api/internal/middleware"
	"sharecvapi/app/model"
	"sharecvapi/common/notice"
	"sharecvapi/common/token"
)

type ServiceContext struct {
	Config   config.Config
	Redis    *redis.Redis
	PermAuth rest.Middleware
	GormDB   *gorm.DB

	AppUserModel         model.AppUserModel
	AppUserInfoModel     model.AppUserInfoModel
	AppUserLocationModel model.AppUserLocationModel
	AppUserFeedbackModel model.AppUserFeedbackModel

	AppConfigModel         model.AppConfigModel
	AppCountryModel        model.AppCountryModel
	AppLanguageModel       model.AppLanguageModel
	AppUserInvitationModel model.AppUserInvitationModel
	AppUserLoginLogModel   model.AppUserLoginLogModel

	AppUserVideoModel            model.AppUserVideoModel
	AppUserPhotoModel            model.AppUserPhotoModel
	AppUserCompanyModel          model.AppUserCompanyModel
	AppUserProjectModel          model.AppUserProjectModel
	AppUserEducationModel        model.AppUserEducationModel
	AppUserVolunteerModel        model.AppUserVolunteerModel
	AppUserExperienceModel       model.AppUserExperienceModel
	AppUserCustomAttributesModel model.AppUserCustomAttributesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	// 自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	//db.AutoMigrate(&models.User{})
	// 初始化JWT
	token.Init(c.JwtAuth.AccessSecret, c.JwtAuth.AccessExpire)
	// 初始化推送
	notice.Init(c.OneSignal.Url, c.OneSignal.AppId, c.OneSignal.RestApiKey)
	// 初始化SMS
	notice.InitTwilio(c.Twilio.AccountSid, c.Twilio.AccountToken, c.Twilio.SmsFrom)

	return &ServiceContext{
		Config:               c,
		Redis:                redisClient,
		GormDB:               db,
		PermAuth:             middleware.NewPermAuthMiddleware(redisClient).Handle,
		AppUserModel:         model.NewAppUserModel(mysqlConn, c.Cache),
		AppUserInfoModel:     model.NewAppUserInfoModel(mysqlConn, c.Cache),
		AppUserLocationModel: model.NewAppUserLocationModel(mysqlConn, c.Cache),
		AppUserFeedbackModel: model.NewAppUserFeedbackModel(mysqlConn, c.Cache),

		AppConfigModel:         model.NewAppConfigModel(mysqlConn, c.Cache),
		AppCountryModel:        model.NewAppCountryModel(mysqlConn, c.Cache),
		AppLanguageModel:       model.NewAppLanguageModel(mysqlConn, c.Cache),
		AppUserInvitationModel: model.NewAppUserInvitationModel(mysqlConn, c.Cache),
		AppUserLoginLogModel:   model.NewAppUserLoginLogModel(mysqlConn, c.Cache),

		AppUserVideoModel:            model.NewAppUserVideoModel(mysqlConn, c.Cache),
		AppUserPhotoModel:            model.NewAppUserPhotoModel(mysqlConn, c.Cache),
		AppUserCompanyModel:          model.NewAppUserCompanyModel(mysqlConn, c.Cache),
		AppUserProjectModel:          model.NewAppUserProjectModel(mysqlConn, c.Cache),
		AppUserEducationModel:        model.NewAppUserEducationModel(mysqlConn, c.Cache),
		AppUserVolunteerModel:        model.NewAppUserVolunteerModel(mysqlConn, c.Cache),
		AppUserExperienceModel:       model.NewAppUserExperienceModel(mysqlConn, c.Cache),
		AppUserCustomAttributesModel: model.NewAppUserCustomAttributesModel(mysqlConn, c.Cache),
	}
}
