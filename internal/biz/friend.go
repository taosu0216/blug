package biz

import (
	v1 "blug/api/blug/v1"
	"context"
)

func (uc *BlugUsecase) CreateNewFriendLinkReq(ctx context.Context, req *v1.CreateNewFriendLinkReq) (*v1.CreateNewFriendLinkResp, error) {
	uc.Log.Infof("CreateNewFriendLinkReq: %v", req)
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
