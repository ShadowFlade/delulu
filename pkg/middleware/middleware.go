package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsSameSite(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
        fmt.Println(c.Request().Header," header")
        fmt.Println(c.Request().Header.Get("Origin")," header2")
		is := c.Request().Header.Get("sec-fetch-site") == "same-origin"
		if is {
			return next(c)
		} else {
			return c.String(http.StatusBadRequest, "suck it")
		}
	}
}
