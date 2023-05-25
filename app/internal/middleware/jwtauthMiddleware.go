package middleware

import (
	"context"
	"git.154896.xyz/zhifou/pkg/jwtx"
	"git.154896.xyz/zhifou/pkg/result"
	"net/http"
)

type JwtAuthMiddleware struct {
	AccessSecret string
}

func NewJwtAuthMiddleware(accessSecret string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		AccessSecret: accessSecret,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ""
		cookie, err := r.Cookie("Authorization")
		header := r.Header.Get("Authorization")
		if err == nil {
			// 有cookie
			id, err = jwtx.ParseToken(m.AccessSecret, cookie.Value)
			if id != "" {
				ctx := r.Context()
				ctx = context.WithValue(ctx, "uid", id)
				next(w, r.WithContext(ctx))
				return
			}
		}
		if header != "" {
			// 有header
			id, err = jwtx.ParseToken(m.AccessSecret, header)
			if id != "" {
				ctx := r.Context()
				ctx = context.WithValue(ctx, "uid", id)
				next(w, r.WithContext(ctx))
				return
			}
		}
		result.HttpResult(r, w, nil, err)
		return
	}
}
