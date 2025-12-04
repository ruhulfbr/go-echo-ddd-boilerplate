package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
)

type UrlDownloaderHandler struct {
	UrlDownloader UrlDownloaderService
}

func NewUrlDownloaderHandler(UrlDownloader UrlDownloaderService) *UrlDownloaderHandler {
	return &UrlDownloaderHandler{UrlDownloader: UrlDownloader}
}

func (h *UrlDownloaderHandler) Download(c echo.Context) error {
	//var registerRequest requests.RegisterRequest
	//if err := c.Bind(&registerRequest); err != nil {
	//	return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request")
	//}
	//
	//if err := registerRequest.Validate(); err != nil {
	//	return err
	//}
	//
	//_, err := h.UrlDownloader.GetUserByEmail(c.Request().Context(), registerRequest.Email)
	//if err == nil {
	//	return responses.ErrorResponse(c, http.StatusConflict, "User already exists")
	//} else if !errors.Is(err, appErrors.ErrUserNotFound) {
	//	return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to check if user exists")
	//}
	//
	//if err := h.UrlDownloader.Register(c.Request().Context(), &registerRequest); err != nil {
	//	return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user")
	//}

	return responses.MessageResponse(c, http.StatusCreated, "Download urls successfully")
}
