package handlers

import (
	"github.com/Rendyfranzz/Study-Case-Chapter-1/presentation"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
)

func GetNews(newsService *service.NewsService) echo.HandlerFunc {
	return func(e echo.Context) error {
		out := newsService.GetNews(e.Request().Context())
		return presentation.Output(e, &out, out.CommonResponse)
	}
}
