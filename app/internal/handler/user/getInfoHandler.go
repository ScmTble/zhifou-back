package user

import (
	"git.154896.xyz/zhifou/pkg/result"
	"net/http"

	"git.154896.xyz/zhifou/app/internal/logic/user"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
