package middleware

import (
	"net/http"
	"todo-app/internal/config"
	"todo-app/internal/libs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthenticateJwt(env *config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			jwtToken, ok := c.Request().Header["Authorization"]
			if !ok {
				return libs.CustomError{
					HTTPCode: http.StatusUnauthorized,
					Message: "Jwt-Token is not found in request header",
				}
			}
			
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(jwtToken[0], claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(env.JwtSecret), nil
			})

			if err != nil {
				return libs.CustomError{
					HTTPCode: http.StatusUnauthorized,
					Message: err.Error(),
				}
			}

			userID, ok := claims["user_id"]
			if !ok {
				return libs.CustomError{
					HTTPCode: http.StatusUnauthorized,
					Message: "Jwt claims does not have user ID",
				}
			}

			c.Set("user_id", int64(userID.(float64)))
			return next(c)
		}
	}
	 
}