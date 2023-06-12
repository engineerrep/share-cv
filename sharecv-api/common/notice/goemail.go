package notice

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"sharecvapi/common/ctxdata"
)

type GoEmailConfig struct {
	Host        string
	Port        int
	Password    string
	FromAddress string
	FromName    string
	AppName     string
}

// GoEmailCode 私人邮箱配置发送验证码
func GoEmailCode(conf GoEmailConfig, language string, emails map[string]string, codeStr string) error {

	var body string
	var subject string
	switch language {
	case ctxdata.LanguageZH:
		subject = conf.AppName + "-验证码"
		body = fmt.Sprintf("<html><body><h2>验证码: </h2><h1>%s</h1></body></html>", codeStr)
	default:
		subject = conf.AppName + "-Code"
		body = fmt.Sprintf("<html><body><h2>Code: </h2><h1>%s</h1></body></html>", codeStr)
		break
	}
	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", conf.FromAddress, conf.FromName)
	// 收件人
	var toEmails []string
	for i, v := range emails {
		toEmails = append(toEmails, m.FormatAddress(i, v))
	}
	m.SetHeader("To", toEmails...)

	// 主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewDialer(conf.Host, conf.Port, conf.FromAddress, conf.Password)

	return d.DialAndSend(m)
}
