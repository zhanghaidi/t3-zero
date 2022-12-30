package ask

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAskEndLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAskEndLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAskEndLogic {
	return &UserAskEndLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAskEndLogic) UserAskEnd(req *types.UserAskEndReq, r *http.Request) error {
	// todo: add your logic here and delete this line

	return nil
}
