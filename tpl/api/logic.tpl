package {{.pkgName}}

import (
	"net/http"
	{{.imports}}
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{if .request}}{{.request}},{{end}}r *http.Request) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
