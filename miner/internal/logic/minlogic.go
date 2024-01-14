package logic

import (
	"context"
	"fmt"

	"github.com/haoshuwei/micro/miner/internal/svc"
	"github.com/haoshuwei/micro/miner/miner"

	"github.com/zeromicro/go-zero/core/logx"
)

type MinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MinLogic {
	return &MinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MinLogic) Min(in *miner.MinerReq) (*miner.MinerResp, error) {
	// todo: add your logic here and delete this line

	var s miner.MinerResp                           // 新增的
	s.Sum = in.A - in.B                             // 新增的
	fmt.Printf("%d - %d = %d\n", in.A, in.B, s.Sum) // 新增的
	return &s, nil                                  // 新增的
}
