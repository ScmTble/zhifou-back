package tag

import (
	"net/http"

	"git.154896.xyz/zhifou/pkg/result"

	"git.154896.xyz/zhifou/app/internal/logic/tag"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag.NewListTagLogic(r.Context(), svcCtx)
		resp, err := l.ListTag(&req)
		result.HttpResult(r, w, resp, err)
	}
}
