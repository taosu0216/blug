package data

import (
	v1 "blug/api/blug/v1"
	"context"
)

func (f *blugRepo) CreateNewFriendLinkInDB(ctx context.Context, req *v1.CreateNewFriendLinkReq) error {
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
