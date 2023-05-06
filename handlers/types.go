package handlers

import "github.com/labstack/echo/v4"

type Handler = func(e echo.Context) error
