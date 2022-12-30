package ask

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAskPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAskPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAskPostLogic {
	return &UserAskPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAskPostLogic) UserAskPost(req *types.UserAskPostReq, r *http.Request) error {
	// todo: add your logic here and delete this line

	return nil
}
