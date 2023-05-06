package handlers

import (
	"net/http"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/presentation"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
)

func Login(authService *service.AuthService) Handler {
	return func(e echo.Context) error {
		in := service.LoginInput{}

		if err := e.Bind(&in); err != nil {
			return err
		}

		out := authService.Login(e.Request().Context(), in)

		return presentation.Output(e, &out, out.CommonResponse)
	}
}

func Register(authService *service.AuthService) Handler {
	return func(e echo.Context) error {
		in := service.RegisterInput{}

		if err := e.Bind(&in); err != nil {
			return err
		}

		out := authService.Register(e.Request().Context(), in)

		return presentation.Output(e, &out, out.CommonResponse)
	}
}

func Logout() Handler {
	return func(e echo.Context) error {
		cookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			HttpOnly: true,
			MaxAge:   -1,
		}

		e.SetCookie(cookie)

		resp := service.CommonResponse{}
		resp.SetMsg(200, "Logout")

		return e.JSON(200, resp)
	}
}
