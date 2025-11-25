package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	errors2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/errors"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	apiResponses "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
)

type authService interface {
	GenerateToken(ctx context.Context, request *requests.LoginRequest) (*apiResponses.LoginResponse, error)
	RefreshToken(ctx context.Context, request *requests.RefreshRequest) (*apiResponses.LoginResponse, error)
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
		return apiResponses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	if err := request.Validate(); err != nil {
		return apiResponses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	response, err := h.authService.GenerateToken(c.Request().Context(), &request)
	switch {
	case errors.Is(err, errors2.ErrUserNotFound), errors.Is(err, errors2.ErrInvalidPassword):
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
	case errors.Is(err, errors2.ErrUserNotFound), errors.Is(err, errors2.ErrInvalidAuthToken):
		return apiResponses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	case err != nil:
		return apiResponses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return apiResponses.Response(c, http.StatusOK, response)
}
