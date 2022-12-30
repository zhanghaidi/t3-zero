package {{.PkgName}}

import (
	"net/http"
	"t3-zero/common/core/response"{{if .HasRequest}}
	"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}{{end}}
		
		l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req,{{end}}r)
		{{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}
	}
}
