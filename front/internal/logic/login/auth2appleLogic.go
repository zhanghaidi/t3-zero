package login

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Auth2appleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuth2appleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Auth2appleLogic {
	return &Auth2appleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Auth2appleLogic) Auth2apple(req *types.Auth2appleReq, r *http.Request) (resp *types.Auth2appleReply, err error) {
	// todo: add your logic here and delete this line

	return
}
