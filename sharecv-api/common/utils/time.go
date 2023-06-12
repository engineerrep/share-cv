package utils

import (
	"sharecvapi/common/globalkey"
	"time"
)

// TimeFormatYMDHMS 格式化 2006-01-02 15:04:05
func TimeFormatYMDHMS(timeUnix int64) string {
	return TimeNow().Format(globalkey.APPDateFormat)
}

// TimeFormatYMD 格式化 2006-01-02
func TimeFormatYMD(timeUnix int64, format string) string {
	return TimeNow().Format(format)
}

// TimeNow 系统时间
func TimeNow() time.Time {
	return time.Now().Local()
}

// TimeNowStr 系统时间
func TimeNowStr() string {
	return TimeNow().Format(globalkey.APPDateFormat)
}
