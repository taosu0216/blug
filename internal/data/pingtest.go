package data

import (
	"blug/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ptRepo struct {
	data *Data
	log  *log.Helper
}

func NewPTRepo(data *Data, logger log.Logger) biz.PingTestRepo {
	return &ptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (pt *ptRepo) Echo(ctx context.Context, p *biz.PingTest) (*biz.PingTest, error) {
	return p, nil
}
