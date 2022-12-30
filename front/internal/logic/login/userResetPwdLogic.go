package login

import (
	"context"
	"net/http"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserResetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserResetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserResetPwdLogic {
	return &UserResetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserResetPwdLogic) UserResetPwd(req *types.ResetPasswordReq, r *http.Request) (resp *types.ResetPasswordReply, err error) {
	// todo: add your logic here and delete this line

	return
}
