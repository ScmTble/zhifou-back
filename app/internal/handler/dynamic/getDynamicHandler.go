package dynamic

import (
	"git.154896.xyz/zhifou/pkg/result"
	"net/http"

	"git.154896.xyz/zhifou/app/internal/logic/dynamic"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDynamicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDynamicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dynamic.NewGetDynamicLogic(r.Context(), svcCtx, r)
		resp, err := l.GetDynamic(&req)
		result.HttpResult(r, w, resp, err)
	}
}
