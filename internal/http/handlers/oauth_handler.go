package handlers

import (
	"context"
	"net/http"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/responses"

	"github.com/labstack/echo/v4"
)

type userAuthenticator interface {
	GoogleOAuth(ctx context.Context, token string) (accessToken string, refreshToken string, exp int64, err error)
}

type OAuthHandler struct {
	userService userAuthenticator
}

func NewOAuthHandler(userService userAuthenticator) *OAuthHandler {
	return &OAuthHandler{userService: userService}
}

func (oa *OAuthHandler) GoogleOAuth(c echo.Context) error {
	var oAuthRequest requests.OAuthRequest

	if err := c.Bind(&oAuthRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	if err := oAuthRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or invalid")
	}

	accessToken, refreshToken, exp, err := oa.userService.GoogleOAuth(c.Request().Context(), oAuthRequest.Token)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to authenticate with Google: "+err.Error())
	}

	res := responses.NewLoginResponse(accessToken, refreshToken, exp)
	return responses.Response(c, http.StatusOK, res)
}
