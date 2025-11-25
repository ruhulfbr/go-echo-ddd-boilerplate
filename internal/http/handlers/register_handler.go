package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	errors2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/errors"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/user"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
)

type userRegisterer interface {
	GetUserByEmail(ctx context.Context, email string) (user.User, error)
	Register(ctx context.Context, request *requests.RegisterRequest) error
}

type RegisterHandler struct {
	userRegisterer userRegisterer
}

func NewRegisterHandler(userRegisterer userRegisterer) *RegisterHandler {
	return &RegisterHandler{userRegisterer: userRegisterer}
}

func (h *RegisterHandler) Register(c echo.Context) error {
	var registerRequest requests.RegisterRequest
	if err := c.Bind(&registerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	}

	if err := registerRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or invalid")
	}

	_, err := h.userRegisterer.GetUserByEmail(c.Request().Context(), registerRequest.Email)
	if err == nil {
		return responses.ErrorResponse(c, http.StatusConflict, "User already exists")
	} else if !errors.Is(err, errors2.ErrUserNotFound) {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to check if user exists")
	}

	if err := h.userRegisterer.Register(c.Request().Context(), &registerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user")
	}

	return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
}
