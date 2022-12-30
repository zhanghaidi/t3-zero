package login

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Auth2qqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuth2qqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Auth2qqLogic {
	return &Auth2qqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Auth2qqLogic) Auth2qq(req *types.Auth2qqReq, r *http.Request) (resp *types.Auth2qqReply, err error) {
	// todo: add your logic here and delete this line

	return
}
