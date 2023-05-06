package presentation

import (
	"net/http"
	"time"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
)

func Output[T any](c echo.Context, out *T, cr service.CommonResponse) error {
	if cr.StatusCode == 0 {
		cr.StatusCode = 200
	}

	if cr.SetAuthToken != `` {
		cookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    cr.SetAuthToken,
			Expires:  time.Now().Add(time.Minute * 2),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}

		c.SetCookie(cookie)
	}

	return c.JSON(cr.StatusCode, out)
}
