package article

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"t3-zero/common/core/response"
	"t3-zero/front/internal/logic/article"
	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"
)

func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewGetArticleLogic(r.Context(), svcCtx)
		resp, err := l.GetArticle(&req, r)
		response.Response(w, resp, err)
	}
}
