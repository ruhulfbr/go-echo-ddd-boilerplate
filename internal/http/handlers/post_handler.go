package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ccoveille/go-safecast"
	"github.com/labstack/echo/v4"
	appErrors "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/errors"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/post"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
)

type PostHandlers struct {
	postService post.Service
}

func NewPostHandlers(postService post.Service) *PostHandlers {
	return &PostHandlers{postService: postService}
}

func (p *PostHandlers) CreatePost(c echo.Context) error {
	authClaims, err := getAuthClaims(c)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	var createPostRequest requests.CreatePostRequest
	if err := c.Bind(&createPostRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := createPostRequest.Validate(); err != nil {
		return err
	}

	post := &models.Post{
		Title:   createPostRequest.Title,
		Content: createPostRequest.Content,
		UserID:  authClaims.ID,
	}

	if err := p.postService.Create(c.Request().Context(), post); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to create post: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "Post successfully created")
}

func (p *PostHandlers) GetPosts(c echo.Context) error {
	posts, err := p.postService.GetPosts(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Failed to get all posts: "+err.Error())
	}

	response := responses.NewPostResponse(posts)
	return responses.Response(c, http.StatusOK, response)
}

func (p *PostHandlers) UpdatePost(c echo.Context) error {
	auth, err := getAuthClaims(c)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	id, err := safecast.ToUint(parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	var updatePostRequest requests.UpdatePostRequest
	if err := c.Bind(&updatePostRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := updatePostRequest.Validate(); err != nil {
		return err
	}

	post, err := p.postService.GetPost(c.Request().Context(), id)
	if errors.Is(err, appErrors.ErrPostNotFound) {
		return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
	} else if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to find post: "+err.Error())
	}

	if post.UserID != auth.ID {
		return responses.ErrorResponse(c, http.StatusForbidden, "Forbidden")
	}

	if err := p.postService.Update(c.Request().Context(), &post, updatePostRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to update post: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusOK, "Post successfully updated")
}

func (p *PostHandlers) DeletePost(c echo.Context) error {
	auth, err := getAuthClaims(c)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	id, err := safecast.ToUint(parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	post, err := p.postService.GetPost(c.Request().Context(), id)
	if errors.Is(err, appErrors.ErrPostNotFound) {
		return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
	} else if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to find post: "+err.Error())
	}

	if post.UserID != auth.ID {
		return responses.ErrorResponse(c, http.StatusForbidden, "Forbidden")
	}

	if err := p.postService.Delete(c.Request().Context(), &post); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete post: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusNoContent, "Post deleted successfully")
}
