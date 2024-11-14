package biz

import (
	v1 "blug/api/blug/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BlugRepo interface {
	CreateNewFriendLinkInDB(ctx context.Context, req *v1.CreateNewFriendLinkReq) error
}

type BlugUsecase struct {
	repo BlugRepo
	log  *log.Helper
}

func NewBlugUsecase(repo BlugRepo, logger log.Logger) *BlugUsecase {
	return &BlugUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BlugUsecase) CreateNewFriendLinkReq(ctx context.Context, req *v1.CreateNewFriendLinkReq) (*v1.CreateNewFriendLinkResp, error) {
	uc.log.Infof("CreateNewFriendLinkReq: %v", req)
	return &v1.CreateNewFriendLinkResp{
			Message: "success",
			Check:   nil,
		}, uc.repo.CreateNewFriendLinkInDB(ctx, &v1.CreateNewFriendLinkReq{
			Title:  req.Title,
			Link:   req.Link,
			Desc:   req.Desc,
			Avatar: req.Avatar,
		})
}
