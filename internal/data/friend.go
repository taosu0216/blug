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
	//start := time.Now()
	//defer func() {
	//	log.Printf("CreateNewFriendLinkReq data took %v", time.Since(start))
	//}()
	//tracer := otel.Tracer("data-tracer")
	//_, span := tracer.Start(ctx, "dataMethod")
	//defer span.End()

	if err != nil {
		f.log.Error(err)
		return err
	}
	return nil
}
