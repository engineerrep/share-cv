package oss

import (
	"context"
	"github.com/aliyun/aliyun-sts-go-sdk/sts"
	"net/http"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssAuthServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssAuthServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssAuthServerLogic {
	return &OssAuthServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssAuthServerLogic) OssAuthServer() (resp *types.OssResp, err error) {
	stsClient := sts.NewClient(l.svcCtx.Config.Ali.AccessKeyID, l.svcCtx.Config.Ali.AccessKeySecret, l.svcCtx.Config.AliOSS.RoleArn, l.svcCtx.Config.AliOSS.RoleSessionName)

	ossResp, err := stsClient.AssumeRole(3600)
	if err != nil {
		return &types.OssResp{
			StatusCode:      http.StatusInternalServerError,
			ErrorCode:       errorx.MapErrMsg(errorx.ServerErrorCode, ctxdata.GetLanguageFromCtx(l.ctx)),
			ErrorMessage:    err.Error(),
			AccessKeyId:     "",
			AccessKeySecret: "",
			SecurityToken:   "",
			Expiration:      "",
		}, nil
	}

	return &types.OssResp{
		StatusCode:      http.StatusOK,
		ErrorCode:       "",
		ErrorMessage:    "",
		AccessKeyId:     ossResp.Credentials.AccessKeyId,
		AccessKeySecret: ossResp.Credentials.AccessKeySecret,
		SecurityToken:   ossResp.Credentials.SecurityToken,
		Expiration:      ossResp.Credentials.Expiration.String(),
	}, nil
}
