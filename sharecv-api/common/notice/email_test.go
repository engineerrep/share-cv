package notice

import (
	"fmt"
	"testing"
)

func TestOnesignalEmail(t *testing.T) {
	Init("https://onesignal.com/api/v1/notifications", "xxxxxx", "xxxxxx")

	var str []string
	str = append(str, "xxxxxx@gmail.com")
	str = append(str, "xxxxxx@gmail.com")
	resCode, err := OnesignalEmailCode("en", str, "Share-CV", "xxxxxx")

	if err != nil {
		fmt.Println("err : ", err.Error())
	}

	if resCode == 200 {
		fmt.Println("邮件发送成功")
	}
}

func TestAliEmailCode(t *testing.T) {
	conf := AliEmailConfig{AccessKeyID: "xxxxxx",
		AccessKeySecret: "xxxxxx",
		AccountName:     "service@xxxxxx.com",
		AddressType:     1,
		Endpoint:        "dm.aliyuncs.com",
	}

	resCode, err := AliEmailCode(conf, "en", "xxxxxx", "xxxxxx@gmail.com")
	if err != nil {
		fmt.Println("err : ", err.Error())
	}

	if resCode == 200 {
		fmt.Println("ali邮件发送成功")
	} else {
		fmt.Println("ali邮件发送失败")
	}
}
