package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/models"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/responses"

	"github.com/labstack/echo/v4"
)

type authService interface {
	GenerateToken(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error)
	RefreshToken(ctx context.Context, request *requests.RefreshRequest) (*responses.LoginResponse, error)
}

type AuthHandler struct {
	authService authService
}

func NewAuthHandler(authService authService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var request requests.LoginRequest
	if err := c.Bind(&request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	response, err := h.authService.GenerateToken(c.Request().Context(), &request)
	switch {
	case errors.Is(err, models.ErrUserNotFound), errors.Is(err, models.ErrInvalidPassword):
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	case err != nil:
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return responses.Response(c, http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var request requests.RefreshRequest
	if err := c.Bind(&request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	response, err := h.authService.RefreshToken(c.Request().Context(), &request)
	switch {
	case errors.Is(err, models.ErrUserNotFound), errors.Is(err, models.ErrInvalidAuthToken):
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	case err != nil:
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return responses.Response(c, http.StatusOK, response)
}
