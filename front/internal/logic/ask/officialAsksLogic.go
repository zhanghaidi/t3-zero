package ask

import (
	"context"
	"net/http"
	"t3-zero/common/model"

	"t3-zero/front/internal/svc"
	"t3-zero/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfficialAsksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOfficialAsksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfficialAsksLogic {
	return &OfficialAsksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OfficialAsksLogic) OfficialAsks(r *http.Request) (resp *types.SimpleList, err error) {
	// todo: add your logic here and delete this line
	db := l.svcCtx.DB

	list := make([]model.Slide, 0)

	if err = db.Model(&model.Slide{}).Find(&list).Error; err != nil {
		return nil, err
	}

	resp = &types.SimpleList{
		List: list,
	}

	return
}
