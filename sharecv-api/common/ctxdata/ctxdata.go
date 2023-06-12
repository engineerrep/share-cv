package ctxdata

import (
	"context"
	"strconv"
)

var (
	// CtxKeyJwtUserUUID 用户UUID
	CtxKeyJwtUserUUID = "UserUUID"

	// CtxKeyJwtUserId 用户ID
	CtxKeyJwtUserId = "UserId"

	// CtxKeyLanguage 从ctx获取用户 Language 语言
	CtxKeyLanguage = "Content-Language"

	// CtxKeyUserIP 从ctx获取用户 UserIP
	CtxKeyUserIP = "UserIP"

	// LanguageEN 英文
	LanguageEN = "en"

	// LanguageZH  中文
	LanguageZH = "zh"
)

func GetUserIdFromCtx(ctx context.Context) int64 {
	userId, _ := ctx.Value(CtxKeyJwtUserId).(string)
	if len(userId) == 0 {
		return 0
	}
	intId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0
	}
	return intId
}

func GetUserUUIDFromCtx(ctx context.Context) string {
	userUUID, _ := ctx.Value(CtxKeyJwtUserUUID).(string)
	if len(userUUID) == 0 {
		return ""
	}
	return userUUID
}

func GetLanguageFromCtx(ctx context.Context) string {
	language, _ := ctx.Value(CtxKeyLanguage).(string)
	if len(language) == 0 {
		return LanguageEN
	}
	return language
}

func GetUserIPFromCtx(ctx context.Context) string {
	realIP, _ := ctx.Value(CtxKeyUserIP).(string)
	return realIP
}
