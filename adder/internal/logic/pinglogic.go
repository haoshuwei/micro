package logic

import (
	"context"
	"fmt"

	"github.com/haoshuwei/micro/adder/adder"
	"github.com/haoshuwei/micro/adder/internal/svc"
	"github.com/haoshuwei/micro/adder/protoc/add"

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

func (l *AddLogic) Add(in *adder.Request) (*adder.Response, error) {
	// todo: add your logic here and delete this line

	var s add.AddResp
	s.Sum = in.A + in.B
	fmt.Printf("%d + %d = %d\n", in.A, in.B, s.Sum)
	//return &add.AddResp{}, nil // 删除的
	return &s, nil
}
