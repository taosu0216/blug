package service

import (

	pb "blug/api/blug/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

