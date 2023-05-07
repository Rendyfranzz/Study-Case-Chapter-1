package service

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/repo"
)

type NewsService struct {
	newsRepo *repo.NewsRepo
}

func NewNewsService(newsRepo *repo.NewsRepo) *NewsService {
	return &NewsService{newsRepo: newsRepo}
}

type GetNewsInput struct {
	CommonRequest
}

type GetNewsOutput struct {
	CommonResponse
	News []model.News `json:"news"`
}

func (n *NewsService) GetNews(ctx context.Context) GetNewsOutput {
	var out GetNewsOutput

	news, err := n.newsRepo.GetNews(ctx)
	if err != nil {
		out.SetMsg(400, err.Error())
		return out
	}
	out.News = news
	out.SetMsg(200, "Success")
	return out
}
