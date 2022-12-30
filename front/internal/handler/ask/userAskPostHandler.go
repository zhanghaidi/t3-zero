package ask

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"t3-zero/common/core/response"
	"t3-zero/front/internal/logic/ask"
	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"
)

func UserAskPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAskPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := ask.NewUserAskPostLogic(r.Context(), svcCtx)
		err := l.UserAskPost(&req, r)
		response.Response(w, nil, err)
	}
}
