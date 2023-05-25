package user

import (
	"net/http"

	"git.154896.xyz/zhifou/pkg/result"

	"git.154896.xyz/zhifou/app/internal/logic/user"
	"git.154896.xyz/zhifou/app/internal/svc"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		result.HttpResult(r, w, resp, err)
	}
}
