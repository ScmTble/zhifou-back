package tool

import (
	"context"
	"errors"
	"git.154896.xyz/zhifou/pkg/jwtx"
	"net/http"
)

func GetIdFromCtx(ctx context.Context) (string, error) {
	value := ctx.Value("uid")
	if value == nil {
		return "", errors.New("get uid from ctx error")
	}
	id, ok := value.(string)
	if !ok {
		return "", errors.New("reflect uid errors")
	}
	return id, nil
}

func TryGetIdFromReq(req *http.Request, accessSecret string) string {
	uid := ""
	cookie, err := req.Cookie("Authorization")
	header := req.Header.Get("Authorization")
	if err == nil {
		// Cookie中有token
		uid, _ = jwtx.ParseToken(accessSecret, cookie.Value)
	}
	if header != "" {
		// Head中有token
		uid, _ = jwtx.ParseToken(accessSecret, header)
	}
	return uid
}
