package ask

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAsksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAsksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAsksLogic {
	return &UserAsksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAsksLogic) UserAsks(r *http.Request) (resp *types.SimpleList, err error) {
	// todo: add your logic here and delete this line

	return
}
