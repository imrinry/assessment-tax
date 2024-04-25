// middleware.go

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// BasicAuth middleware

// Middleware function to configure basic authentication
func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		username, password, ok := ctx.Request().BasicAuth()
		if !ok || !basicAuth(username, password, ctx) {
			ctx.Response().Header().Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(ctx)
	}
}

func basicAuth(username, password string, ctx echo.Context) bool {
	return username == "adminTax" && password == "admin!"
}
