package notice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sharecvapi/common/ctxdata"
	"strings"
)

var onesignalUrl string
var restAPIKey string
var appId string

type PushContents struct {
	En string `json:"en"`
	Zh string `json:"zh"`
}

type PayloadPush struct {
	AppId                     string       `json:"app_id"`
	IncludeExternalUserIds    []string     `json:"include_external_user_ids"`
	ChannelForExternalUserIds string       `json:"channel_for_external_user_ids"`
	Contents                  PushContents `json:"contents"`
}
type PayloadEmail struct {
	AppId              string   `json:"app_id"`
	IncludeEmailTokens []string `json:"include_email_tokens"`
	EmailFromName      string   `json:"email_from_name"`
	EmailSubject       string   `json:"email_subject"`
	EmailBody          string   `json:"email_body"`
}
type PayloadSMS struct {
	AppId               string       `json:"app_id"`
	IncludePhoneNumbers []string     `json:"include_phone_numbers"`
	SmsMediaUrls        []string     `json:"sms_media_urls"`
	Contents            PushContents `json:"contents"`
	Name                string       `json:"name"`
	SmsFrom             string       `json:"sms_from"`
}

func Init(url string, id string, key string) {
	appId = id
	restAPIKey = key
	onesignalUrl = url
}

// OnesignalPush 发送推送到指定的用户群体
func OnesignalPush(language string, userIds []string, content string) (int, error) {
	contents := PushContents{
		En: content,
	}
	switch language {
	case ctxdata.LanguageZH:
		contents.Zh = content
	default:
		break
	}
	payload := PayloadPush{
		AppId:                     appId,
		IncludeExternalUserIds:    userIds,
		ChannelForExternalUserIds: "push",
		Contents:                  contents,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	return requestOneSignal(string(jsonPayload))
}

// OnesignalEmailCode 发送邮箱验证码
func OnesignalEmailCode(language string, userEmails []string, appName string, codeStr string) (int, error) {
	var body string
	var subject string
	switch language {
	case ctxdata.LanguageZH:
		subject = appName + "-验证码"
		body = fmt.Sprintf("<html><body><h2>验证码: </h2><h1>%s</h1></body></html>", codeStr)
	default:
		subject = appName + "-Code"
		body = fmt.Sprintf("<html><body><h2>Code: </h2><h1>%s</h1></body></html>", codeStr)
		break
	}
	payload := PayloadEmail{
		AppId:              appId,
		IncludeEmailTokens: userEmails,
		EmailSubject:       subject,
		EmailFromName:      appName,
		EmailBody:          body,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	return requestOneSignal(string(jsonPayload))
}

// OnesignalSMS 发送手机验证码
func OnesignalSMS(language string, appName string, userPhones []string, smsFrom string, code string) (int, error) {
	contents := PushContents{
		En: fmt.Sprintf("%s Code: %s", appName, code),
	}
	switch language {
	case ctxdata.LanguageZH:
		contents.Zh = fmt.Sprintf("%s 验证码: %s", appName, code)
	default:
		break
	}
	payload := PayloadSMS{
		AppId:               appId,
		IncludePhoneNumbers: userPhones,
		Contents:            contents,
		Name:                appName,
		SmsFrom:             smsFrom,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	return requestOneSignal(string(jsonPayload))
}

// 发送消息到OneSignal
func requestOneSignal(content string) (int, error) {
	payload := strings.NewReader(content)

	req, _ := http.NewRequest("POST", onesignalUrl, payload)

	req.Header.Add("Authorization", "Basic "+restAPIKey)
	req.Header.Add("content-type", "application/json; charset=utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	//defer res.Body.Close()
	//body, err := io.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	return res.StatusCode, err
}
