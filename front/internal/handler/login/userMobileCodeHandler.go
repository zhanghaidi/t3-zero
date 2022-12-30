package login

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"t3-zero/common/core/response"
	"t3-zero/front/internal/logic/login"
	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"
)

func UserMobileCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMobileCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := login.NewUserMobileCodeLogic(r.Context(), svcCtx)
		err := l.UserMobileCode(&req, r)
		response.Response(w, nil, err)
	}
}
