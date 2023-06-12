package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net"
	"net/http"
	"sharecvapi/common/ctxdata"
	"strings"
)

type PermAuthMiddleware struct {
	Redis *redis.Redis
}

func NewPermAuthMiddleware(r *redis.Redis) *PermAuthMiddleware {
	return &PermAuthMiddleware{
		Redis: r,
	}
}

func (m *PermAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		language := r.Header.Get(ctxdata.CtxKeyLanguage)
		ctx = context.WithValue(ctx, ctxdata.CtxKeyLanguage, language)

		ip := r.Header.Get("X-Real-IP")
		if net.ParseIP(ip) != nil {
			ctx = context.WithValue(ctx, ctxdata.CtxKeyUserIP, ip)
		} else {
			ip = r.Header.Get("X-Forward-For")
			var found = false
			for _, i := range strings.Split(ip, ",") {
				if net.ParseIP(i) != nil {
					found = true
					ctx = context.WithValue(ctx, ctxdata.CtxKeyUserIP, ip)
				}
			}
			if !found {
				ip, _, _ = net.SplitHostPort(r.RemoteAddr)
				if net.ParseIP(ip) != nil {
					ctx = context.WithValue(ctx, ctxdata.CtxKeyUserIP, ip)
				}
			}
		}

		next(w, r.WithContext(ctx))
	}
}
