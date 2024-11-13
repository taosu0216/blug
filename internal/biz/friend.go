package biz

import (
	v1 "blug/api/blug/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FriendRepo interface {
	CreateNewFriendLinkInDB(ctx context.Context, req *v1.CreateNewFriendLinkReq) error
}

type FriendUsecase struct {
	repo FriendRepo
	log  *log.Helper
}

func NewFriendUsecase(repo FriendRepo, logger log.Logger) *FriendUsecase {
	return &FriendUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FriendUsecase) CreateNewFriendLinkReq(ctx context.Context, req *v1.CreateNewFriendLinkReq) (*v1.CreateNewFriendLinkResp, error) {
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
