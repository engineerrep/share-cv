package verify

import (
	"context"
	"github.com/GianOrtiz/apple-auth-go"
	"google.golang.org/api/oauth2/v2"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"time"
)

const appleId = "com.yisheng.app"
const appleTeamId = "xxxxxx"
const appleKeyId = "xxxxxx"
const appleKeyPath = "./key/AuthKey_N6976RK64F.p8"

func AppleVerifyCode(ctx context.Context, authorizationCode string) (error, *apple.AppleUser) {

	appleAuth, err := apple.New(appleId, appleTeamId, appleKeyId, appleKeyPath)
	if err != nil {
		return err, nil
	}

	tokenResponse, err := appleAuth.ValidateCode(authorizationCode)
	if err != nil {
		return err, nil
	}

	user, err := apple.GetUserInfoFromIDToken(tokenResponse.IDToken)
	if err != nil {
		return err, nil
	}

	return nil, user
}

func GoogleVerifyToken(ctx context.Context, idToken string, accessToken string) (error, *oauth2.Tokeninfo) {
	service, err := oauth2.NewService(ctx)
	if err != nil {
		return err, nil
	}

	token := service.Tokeninfo()
	token.IdToken(idToken)
	token.AccessToken(accessToken)
	info, err := token.Do()
	if err != nil {
		return err, nil
	}

	return nil, info
}

func GoogleVerifyReCaptchaV3(token string, key string) error {
	reCAPTCHA, err := recaptcha.NewReCAPTCHA(key, recaptcha.V3, time.Second*10)
	if err != nil {
		return err
	}

	err = reCAPTCHA.Verify(token)
	if err != nil {
		return err
	}

	return nil
}
