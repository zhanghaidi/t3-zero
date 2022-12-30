package ask

import (
	"net/http"
	"t3-zero/common/core/response"
	"t3-zero/front/internal/logic/ask"
	"t3-zero/front/internal/svc"
)

func UserAsksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := ask.NewUserAsksLogic(r.Context(), svcCtx)
		resp, err := l.UserAsks(r)
		response.Response(w, resp, err)
	}
}
