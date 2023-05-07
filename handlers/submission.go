package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/presentation"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
)

func CreateSubmissions(SubmissionService *service.SubmissionService) echo.HandlerFunc {
	return func(e echo.Context) error {
		in := service.SubmissionRegisterInput{}

		if err := e.Bind(&in); err != nil {
			return err
		}
		fmt.Println(in)
		out := SubmissionService.RegisterSubmission(e.Request().Context(), in)

		return presentation.Output(e, &out, out.CommonResponse)
	}
}

func EditSubmission(SubmissionService *service.SubmissionService) echo.HandlerFunc {
	return func(e echo.Context) error {
		key := e.Param("nik")
		in := service.SubmissionEditInput{}
		json.NewDecoder(e.Request().Body).Decode(&in)
		if key == "" {
			var out service.CommonResponse
			out.SetMsg(400, "nik is required")
			return e.JSON(400, out)
		}

		out := SubmissionService.EditSubmission(e.Request().Context(), in, service.SubmissionEditKey{Search: key})
		return presentation.Output(e, &out, out.CommonResponse)
	}
}
