package delivery

import (
	"errors"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/post"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type PostHttpDelivery struct {
	api     *echo.Group
	interop post.PostInterop
}

func (p PostHttpDelivery) List(c echo.Context) error {
	query := &common.QueryOpts{}
	pageStr := c.QueryParam("page")
	if pageStr == "" {
		return c.JSON(http.StatusBadRequest, "page is empty")
	}
	sizeStr := c.QueryParam("size")
	if sizeStr == "" {
		return c.JSON(http.StatusBadRequest, "size is empty")
	}
	if pageStr != "" {
		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "page is not a number")
		}
		query.Page = int(page)
	}
	if sizeStr != "" {
		size, err := strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "size is not a number")
		}
		query.Size = int(size)
	}
	_ = c.Bind(query)
	post, err := p.interop.List(c.Request().Context(), query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (p PostHttpDelivery) Create(c echo.Context) error {
	postData := &post.Post{}
	if err := c.Bind(postData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := p.interop.Create(c.Request().Context(), token, postData)

	if err != nil {
		if errors.Is(err, post.ErrPostRequiredContent) || errors.Is(err, post.ErrPostRequiredPhoto) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, post.ErrPostNotCreated) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusCreated, postData)
}

func NewPostHttpDelivery(api *echo.Group, interop post.PostInterop) *PostHttpDelivery {
	handler := &PostHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	api.GET("/all", handler.List)
	return handler
}