package oss

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/mmdb"
	"sharecvapi/common/utils"
	"strings"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/logx"
)

type OssUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssUploadLogic {
	return &OssUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssUploadLogic) OssUpload(r *http.Request) (resp *types.OssUploadResp, err error) {

	_, fileHeader, _ := r.FormFile("file")
	if fileHeader == nil {
		return nil, errorx.NewCodeError(errorx.ParamErrorCode)
	}

	file, _ := fileHeader.Open()
	if file == nil {
		return nil, errorx.NewCodeError(errorx.ParamErrorCode)
	}

	bucketName, bucketEndpoint := GetBucketWithIp(l.ctx, l.svcCtx)
	client, err := alioss.New(bucketEndpoint, l.svcCtx.Config.Ali.AccessKeyID, l.svcCtx.Config.Ali.AccessKeySecret)
	if err != nil {
		l.Logger.Errorf("create client error: %v \n\n", err)
		return nil, errorx.NewCodeError(errorx.UploadErrorCode)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		l.Logger.Errorf("Get Bucket error: %v \n\n", err)
		return nil, errorx.NewCodeError(errorx.UploadErrorCode)
	}

	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	files := strings.Split(fileHeader.Filename, ".")
	objectKey := fmt.Sprintf("share/web/%d/%s.%s", userId, utils.MD5(fileHeader.Filename), files[len(files)-1])
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		l.Logger.Errorf("Put Object error: %v \n\n", err)
		return &types.OssUploadResp{FilePath: ""}, nil
	}

	return &types.OssUploadResp{FilePath: "/" + objectKey}, nil
}

// GetBucketWithIp 根据用户IP判断使用那个储存桶
func GetBucketWithIp(ctx context.Context, svcCtx *svc.ServiceContext) (string, string) {

	if svcCtx.Config.Mode == "dev" || svcCtx.Config.Mode == "pre" {
		return svcCtx.Config.AliOSS.BucketNameHK, svcCtx.Config.AliOSS.BucketEndpointHK
	}

	var realIP = ctxdata.GetUserIPFromCtx(ctx)
	oceania, country, err := mmdb.SearchContinent(realIP)
	if err != nil {
		return svcCtx.Config.AliOSS.BucketNameUS, svcCtx.Config.AliOSS.BucketEndpointUS
	}

	// 判断是否是中国IP
	if country == svcCtx.Config.Region.China {
		return svcCtx.Config.AliOSS.BucketNameAP, svcCtx.Config.AliOSS.BucketEndpointAP
	}

	// 亚洲国家使用新加坡服务器
	if utils.StringIn(oceania, svcCtx.Config.Region.Asia) {
		return svcCtx.Config.AliOSS.BucketNameAP, svcCtx.Config.AliOSS.BucketEndpointAP
	}

	return svcCtx.Config.AliOSS.BucketNameUS, svcCtx.Config.AliOSS.BucketEndpointUS
}

// UploadWithImageUrl 上传图片到OSS并返回图片地址
func UploadWithImageUrl(ctx context.Context, svcCtx *svc.ServiceContext, url string) (string, error) {
	// 从URL下载图片
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取图片 上传oss
	bucketName, bucketEndpoint := GetBucketWithIp(ctx, svcCtx)
	client, err := alioss.New(bucketEndpoint, svcCtx.Config.Ali.AccessKeyID, svcCtx.Config.Ali.AccessKeySecret)
	if err != nil {
		return "", errorx.NewCodeError(errorx.UploadErrorCode)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", errorx.NewCodeError(errorx.UploadErrorCode)
	}

	img, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errorx.NewCodeError(errorx.UploadErrorCode)
	}

	var reader io.Reader = bytes.NewReader(img)
	userId := ctxdata.GetUserIdFromCtx(ctx)
	objectKey := fmt.Sprintf("share/web/%d/%s.%s", userId, utils.MD5(url), "jpg")
	err = bucket.PutObject(objectKey, reader)
	if err != nil {
		return "", err
	}

	return "/" + objectKey, nil
}
