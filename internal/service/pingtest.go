package service

import (
	"blug/internal/biz"
	"context"
	"log"

	pb "blug/api/blug/v1"
)

type PingTestService struct {
	pb.UnimplementedPingTestServer

	uc *biz.PingTestUsecase
}

func NewPingTestService(uc *biz.PingTestUsecase) *PingTestService {
	return &PingTestService{uc: uc}
}

func (s *PingTestService) SayPing(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	sp, err := s.uc.EchoTest(ctx, &biz.PingTest{PTecho: req.Name})
	log.Println("11111111")
	if err != nil {
		return nil, err
	}
	return &pb.PingResp{Message: sp.PTecho}, nil
}
