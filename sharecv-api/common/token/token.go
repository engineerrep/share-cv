package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"sharecvapi/common/ctxdata"
	"strconv"
)

var accessSecret string
var accessExpire int64

func Init(secret string, expire int64) {
	accessSecret = secret
	accessExpire = expire
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return []byte(accessSecret), nil
}

// GenToken 生成JWT
func GenToken(userID int64, userUUID string, iat int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + accessExpire
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserUUID] = userUUID
	claims[ctxdata.CtxKeyJwtUserId] = strconv.FormatInt(userID, 10)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(accessSecret))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (int64, string, error) {
	// 解析token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return 0, "", err
	}
	// 校验token
	if token.Valid {
		userUUID := claims[ctxdata.CtxKeyJwtUserUUID].(string)
		userId := claims[ctxdata.CtxKeyJwtUserId].(string)
		userIntId, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			return 0, "", err
		}
		return userIntId, userUUID, nil
	}
	return 0, "", errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(token string, now int64) (newToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(token, keyFunc); err != nil {
		return "", err
	}
	userId, userUUID, err := ParseToken(token)
	v, _ := err.(*jwt.ValidationError)
	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {

		newToken, err = GenToken(userId, userUUID, now)
		return newToken, err
	}
	return "", err
}
