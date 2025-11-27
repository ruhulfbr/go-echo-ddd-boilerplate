package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/token"
)

func EchoJWTMiddleware(signingKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(signingKey),
	})
}
