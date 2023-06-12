package notice

import (
	"fmt"
	"testing"
)

func TestOnesignalSMS(t *testing.T) {
	Init("https://onesignal.com/api/v1/notifications", "xxxxxx", "xxxxxx")

	var phones []string
	phones = append(phones, "+xxxxxx")
	resCode, err := OnesignalSMS("en", "ShareCV", phones, "+xxxxxx", "xxxxxx")

	if err != nil {
		fmt.Println("err : ", err.Error())
	}

	if resCode == 200 {
		fmt.Println("短信发送成功")
	}
}

func TestTwilioSMS(t *testing.T) {
	InitTwilio("xxxxxx", "xxxxxx", "+xxxxxx")

	//phone := "+xxxxxx"
	phone := "+xxxxxx"
	resCode, err := SendTwilioSMS("en", "xxxxxx", phone)

	if err != nil {
		fmt.Println("err : ", err.Error())
	}

	if resCode == 200 {
		fmt.Println("短信发送成功")
	}
}
