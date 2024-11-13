package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type PingTest struct {
	PTecho string
}

type PingTestRepo interface {
	Echo(context.Context, *PingTest) (*PingTest, error)
}

type PingTestUsecase struct {
	repo PingTestRepo
	log  *log.Helper
}

func NewPingTestUsecase(repo PingTestRepo, logger log.Logger) *PingTestUsecase {
	return &PingTestUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PingTestUsecase) EchoTest(ctx context.Context, pt *PingTest) (*PingTest, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", pt.PTecho)
	return uc.repo.Echo(ctx, pt)
}
