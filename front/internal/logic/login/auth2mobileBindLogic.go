package login

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Auth2mobileBindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuth2mobileBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Auth2mobileBindLogic {
	return &Auth2mobileBindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Auth2mobileBindLogic) Auth2mobileBind(req *types.Auth2mobileBindReq, r *http.Request) (resp *types.Auth2mobileBindReply, err error) {
	// todo: add your logic here and delete this line

	return
}
