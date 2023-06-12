package notice

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"sharecvapi/common/ctxdata"
)

var accountSID string
var accountToken string
var from string

func InitTwilio(sid string, token string, phone string) {
	accountSID = sid
	accountToken = token
	from = phone
}

// SendTwilioSMS 发送短信
func SendTwilioSMS(language string, code string, to string) (int, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: accountToken,
	})
	var body = fmt.Sprintf("Code: %s", code)
	switch language {
	case ctxdata.LanguageZH:
		body = fmt.Sprintf("验证码: %s", code)
	default:
		break
	}
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return 0, err
	}
	return 200, nil
}
