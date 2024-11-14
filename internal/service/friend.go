package service

import (
	pb "blug/api/blug/v1"
	"blug/internal/biz"
	"blug/internal/pkg"
	"context"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"os"
)

func NewFriendService(uc *biz.BlugUsecase) *BlugService {
	return &BlugService{uc: uc}
}

func (s *BlugService) CreateNewFriendLink(ctx context.Context, req *pb.CreateNewFriendLinkReq) (*pb.CreateNewFriendLinkResp, error) {
	res, err := s.uc.CreateNewFriendLinkReq(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateNewFriendLinkResp{Message: res.Message, Check: res.Check}, nil
}

func (s *BlugService) Upload(ctx transporthttp.Context) error {
	file, header, err := ctx.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// 修改文件名并创建保存图片
	imageFile, err := os.Create(pkg.GetPostsLocation() + "band_" + header.Filename)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	// 将文件内容复制到保存的文件中
	_, err = io.Copy(imageFile, file)
	if err != nil {
		return err
	}
	return nil
}
