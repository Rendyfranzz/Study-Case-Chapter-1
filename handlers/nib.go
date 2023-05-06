package handlers

import (
	"github.com/Rendyfranzz/Study-Case-Chapter-1/presentation"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
)

func GetNIB(nibService *service.NIBService) echo.HandlerFunc {
	return func(e echo.Context) error {
		id := e.Param("id")
		if id == "" {
			var out service.CommonResponse
			out.SetMsg(400, "id is required")
			return e.JSON(400, out)
		}

		out := nibService.GetNIB(e.Request().Context(), service.GetNIBInput{ID: id})
		return presentation.Output(e, &out, out.CommonResponse)
	}
}

func CreateNIB(nibService *service.NIBService) echo.HandlerFunc {
	return func(e echo.Context) error {
		in := service.NIBRegisterInput{}

		if err := e.Bind(&in); err != nil {
			return err
		}

		out := nibService.RegisterNIB(e.Request().Context(), in)

		return presentation.Output(e, &out, out.CommonResponse)
	}
}
