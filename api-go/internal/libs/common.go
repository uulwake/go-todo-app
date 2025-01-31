package libs

import "github.com/labstack/echo/v4"

func GetRequestID(ctx echo.Context) string {
	return ctx.Response().Header().Get(echo.HeaderXRequestID)
}