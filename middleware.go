package main

import (
	"github.com/Rendyfranzz/Study-Case-Chapter-1/pkg"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(validator pkg.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(`token`)
			if err != nil {
				return err
			}

			token := cookie.Value

			if err := validator.Validate(token); err != nil {
				resp := service.CommonResponse{}
				resp.SetMsg(400, err.Error())

				return c.JSON(400, resp)
			}

			return next(c)
		}
	}
}
