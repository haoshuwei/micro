package logic

import (
	"context"

	"github.com/haoshuwei/micro/adder/internal/svc"
	"github.com/haoshuwei/micro/adder/protoc/adder"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *adder.AddReq) (*adder.AddResp, error) {
	// todo: add your logic here and delete this line

	return &adder.AddResp{}, nil
}
