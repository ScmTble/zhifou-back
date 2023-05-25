package tag

import (
	"git.154896.xyz/zhifou/pkg/result"
	"net/http"

	"git.154896.xyz/zhifou/app/internal/logic/tag"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag.NewQueryTagLogic(r.Context(), svcCtx)
		resp, err := l.QueryTag(&req)
		result.HttpResult(r, w, resp, err)
	}
}
