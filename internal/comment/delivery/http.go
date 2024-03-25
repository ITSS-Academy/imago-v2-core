package delivery

import (
	"errors"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/comment"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CommentHttpDelivery struct {
	api     *echo.Group
	interop comment.CommentInterop
}

func (c CommentHttpDelivery) CreateComment(e echo.Context) error {
	commentData := &comment.Comment{}
	if err := e.Bind(commentData); err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	token := e.Request().Header.Get("Authorization")
	err := c.interop.CreateComment(e.Request().Context(), token, commentData)
	if err != nil {
		if errors.Is(err, comment.ErrCommentContentEmpty) {
			return e.JSON(http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, comment.ErrCommentNotCreated) {
			return e.JSON(http.StatusInternalServerError, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.JSON(http.StatusCreated, commentData)
}

func (c CommentHttpDelivery) GetCommentById(e echo.Context) error {
	id := e.QueryParam("id")
	token := e.Request().Header.Get("Authorization")
	commentData, err := c.interop.GetCommentById(e.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, comment.ErrCommentNotFound) {
			return e.JSON(http.StatusNotFound, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.JSON(http.StatusOK, commentData)
}

func (c CommentHttpDelivery) GetCommentByPostId(e echo.Context) error {
	postId := e.QueryParam("post_id")
	token := e.Request().Header.Get("Authorization")
	query := &common.QueryOpts{}
	pageStr := e.QueryParam("page")
	if pageStr == "" {
		return e.JSON(http.StatusBadRequest, "page is empty")
	}
	sizeStr := e.QueryParam("size")
	if sizeStr == "" {
		return e.JSON(http.StatusBadRequest, "size is empty")
	}
	if pageStr != "" {
		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "page is not a number")
		}
		query.Page = int(page)
	}
	if sizeStr != "" {
		size, err := strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "size is not a number")
		}
		query.Size = int(size)
	}
	commentData, err := c.interop.GetCommentByPostId(e.Request().Context(), token, postId, query)
	if err != nil {
		if errors.Is(err, comment.ErrCommentNotFound) {
			return e.JSON(http.StatusNotFound, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.JSON(http.StatusOK, commentData)
}

func (c CommentHttpDelivery) GetComment(e echo.Context) error {
	token := e.Request().Header.Get("Authorization")
	query := &common.QueryOpts{}
	pageStr := e.QueryParam("page")
	if pageStr == "" {
		return e.JSON(http.StatusBadRequest, "page is empty")
	}
	sizeStr := e.QueryParam("size")
	if sizeStr == "" {
		return e.JSON(http.StatusBadRequest, "size is empty")
	}
	if pageStr != "" {
		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "page is not a number")
		}
		query.Page = int(page)
	}
	if sizeStr != "" {
		size, err := strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "size is not a number")
		}
		query.Size = int(size)
	}
	commentData, err := c.interop.GetComment(e.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, comment.ErrCommentNotFound) {
			return e.JSON(http.StatusNotFound, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.JSON(http.StatusOK, commentData)
}

func (c CommentHttpDelivery) UpdateComment(e echo.Context) error {
	id := e.QueryParam("id")
	commentData := &comment.Comment{}
	if err := e.Bind(commentData); err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	token := e.Request().Header.Get("Authorization")
	err := c.interop.UpdateComment(e.Request().Context(), token, id, commentData)
	if err != nil {
		if errors.Is(err, comment.ErrCommentIdEmpty) || errors.Is(err, comment.ErrCommentContentEmpty) {
			return e.JSON(http.StatusBadRequest, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.JSON(http.StatusOK, commentData)
}

func (c CommentHttpDelivery) DeleteComment(e echo.Context) error {
	id := e.QueryParam("id")
	token := e.Request().Header.Get("Authorization")
	err := c.interop.DeleteComment(e.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, comment.ErrCommentIdEmpty) {
			return e.JSON(http.StatusBadRequest, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.NoContent(http.StatusOK)
}

func NewCommentHttpDelivery(api *echo.Group, interop comment.CommentInterop) *CommentHttpDelivery {
	handler := &CommentHttpDelivery{api: api, interop: interop}
	api.POST("", handler.CreateComment)
	api.GET("/id", handler.GetCommentById)
	api.GET("/postId", handler.GetCommentByPostId)
	api.GET("/all", handler.GetComment)
	api.PUT("", handler.UpdateComment)
	api.DELETE("", handler.DeleteComment)
	return handler
}
