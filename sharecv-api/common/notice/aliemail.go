package notice

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"sharecvapi/common/ctxdata"
)

type AliEmailConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	AccountName     string
	AddressType     int32
	Endpoint        string
	AppName         string
}

// createClient 创建阿里操作对象
func createClient(conf AliEmailConfig) (_result *dm20151123.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     &conf.AccessKeyID,
		AccessKeySecret: &conf.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = &conf.Endpoint
	_result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(config)
	return _result, _err
}

// AliEmailCode 阿里云邮件推送
func AliEmailCode(conf AliEmailConfig, language string, codeStr string, email string) (int32, error) {
	// 发送邮箱验证码
	var body string
	var subject string
	switch language {
	case ctxdata.LanguageZH:
		subject = conf.AppName + "-验证码"
		body = fmt.Sprintf("<html><body><h4>验证码:</h4><h1>%s</h1></body></html>", codeStr)
	default:
		subject = conf.AppName + "-Code"
		body = fmt.Sprintf("<html><body><h4>Code:</h4><h1>%s</h1></body></html>", codeStr)
	}

	client, err := createClient(conf)
	if err != nil {
		return 0, err
	}

	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String(conf.AccountName),
		AddressType:    tea.Int32(conf.AddressType),
		ReplyToAddress: tea.Bool(true),
		ToAddress:      tea.String(email),
		Subject:        tea.String(subject),
		HtmlBody:       tea.String(body),
	}

	runtime := &util.RuntimeOptions{}
	tryRes, tryErr := func() (_result *dm20151123.SingleSendMailResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		res, err := client.SingleSendMailWithOptions(singleSendMailRequest, runtime)
		if err != nil {
			return nil, err
		}

		return res, nil
	}()

	if tryErr != nil {
		return 0, tryErr
	}

	return *tryRes.StatusCode, tryErr
}
