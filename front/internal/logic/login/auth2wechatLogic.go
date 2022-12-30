package login

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Auth2wechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuth2wechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Auth2wechatLogic {
	return &Auth2wechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Auth2wechatLogic) Auth2wechat(req *types.Auth2wechatReq, r *http.Request) (resp *types.Auth2wechatReply, err error) {
	// todo: add your logic here and delete this line

	return
}
