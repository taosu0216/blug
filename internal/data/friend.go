package data

import (
	v1 "blug/api/blug/v1"
	"blug/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type friendRepo struct {
	data *Data
	log  *log.Helper
}

func NewFriendRepo(data *Data, logger log.Logger) biz.FriendRepo {
	return &friendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f *friendRepo) CreateNewFriendLinkInDB(ctx context.Context, req *v1.CreateNewFriendLinkReq) error {
	err := f.data.DB.Friend.Create().
		SetTitle(req.Title).
		SetLink(req.Link).
		SetDesc(req.Desc).
		SetAvatar(req.Avatar).
		Exec(ctx)

	if err != nil {
		f.log.Errorf("CreateNewFriendLinkInDB error: %v", err)
		return err
	}
	return nil
}
