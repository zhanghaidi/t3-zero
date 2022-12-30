package article

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryArticlesLogic {
	return &CategoryArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryArticlesLogic) CategoryArticles(req *types.CategoryArticlesReq, r *http.Request) (resp *types.PageListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
