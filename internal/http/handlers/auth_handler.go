package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	appErrors "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/errors"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	apiResponses "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/auth"
)

type AuthHandler struct {
	authService auth.AuthServiceInterface
}

func NewAuthHandler(authService auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var request requests.LoginRequest
	if err := c.Bind(&request); err != nil {
		return apiResponses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	if err := request.Validate(); err != nil {
		return err
	}

	response, err := h.authService.GenerateToken(c.Request().Context(), &request)
	switch {
	case errors.Is(err, appErrors.ErrUserNotFound), errors.Is(err, appErrors.ErrInvalidPassword):
		return apiResponses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	case err != nil:
		return apiResponses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return apiResponses.Response(c, http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var request requests.RefreshRequest
	if err := c.Bind(&request); err != nil {
		return apiResponses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	response, err := h.authService.RefreshToken(c.Request().Context(), &request)
	switch {
	case errors.Is(err, appErrors.ErrUserNotFound), errors.Is(err, appErrors.ErrInvalidAuthToken):
		return apiResponses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	case err != nil:
		return apiResponses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return apiResponses.Response(c, http.StatusOK, response)
}
