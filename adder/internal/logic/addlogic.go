package logic

import (
	"context"
	"fmt"

	"github.com/haoshuwei/micro/adder/adder"
	"github.com/haoshuwei/micro/adder/internal/svc"

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
	var s adder.AddResp                             // 新增的
	s.Sum = in.A + in.B                             // 新增的
	fmt.Printf("%d + %d = %d\n", in.A, in.B, s.Sum) // 新增的
	//return &add.AddResp{}, nil // 删除的
	return &s, nil // 新增的
}
