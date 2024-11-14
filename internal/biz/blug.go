package biz

import (
	v1 "blug/api/blug/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BlugRepo interface {
	// friend service
	CreateNewFriendLinkInDB(ctx context.Context, req *v1.CreateNewFriendLinkReq) error
}

type BlugUsecase struct {
	repo BlugRepo
	Log  *log.Helper
}

func NewBlugUsecase(repo BlugRepo, logger log.Logger) *BlugUsecase {
	return &BlugUsecase{repo: repo, Log: log.NewHelper(logger)}
}
