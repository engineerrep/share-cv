package utils

import (
	"bytes"
	"container/list"
	"crypto/rand"
	"fmt"
	"io"
	rand2 "math/rand"
	"time"
)

// GetUuid 获取UUID
func GetUuid() string {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// RandString 随机获取字符
func RandString(codeLen int) string {
	// 1. 定义原始字符串
	rawStr := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	// 随机从中获取
	rand2.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand2.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String()
}

var baseStr string = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
var base []byte = []byte(baseStr)

// InvitationCode 根据用户ID生成邀请码
func InvitationCode(n int64, lens int) string {
	quotient := TimeNow().Unix() + n
	mod := int64(0)
	l := list.New()
	for quotient != 0 {
		mod = quotient % 34
		quotient = quotient / 34
		l.PushFront(base[int(mod)])
	}
	listLen := l.Len()
	if listLen >= lens {
		res := make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
		return string(res)
	} else {
		res := make([]byte, 0, lens)
		for i := 0; i < lens; i++ {
			if i < lens-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}
		}
		return string(res)
	}
}
