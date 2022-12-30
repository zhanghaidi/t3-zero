package login

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMobileCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMobileCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMobileCodeLogic {
	return &UserMobileCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMobileCodeLogic) UserMobileCode(req *types.GetMobileCodeReq, r *http.Request) error {
	// todo: add your logic here and delete this line

	return nil
}
