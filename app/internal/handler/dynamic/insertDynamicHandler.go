package dynamic

import (
	"net/http"

	"git.154896.xyz/zhifou/pkg/result"

	"git.154896.xyz/zhifou/app/internal/logic/dynamic"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InsertDynamicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertPostReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpResult(r, w, nil, err)
			return
		}

		l := dynamic.NewInsertDynamicLogic(r.Context(), svcCtx)
		resp, err := l.InsertDynamic(&req)
		result.HttpResult(r, w, resp, err)
	}
}
