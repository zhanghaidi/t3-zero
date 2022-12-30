package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"t3-zero/backend/internal/logic"
	"t3-zero/backend/internal/svc"
	"t3-zero/backend/internal/types"
)

func BackendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBackendLogic(r.Context(), svcCtx)
		resp, err := l.Backend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
