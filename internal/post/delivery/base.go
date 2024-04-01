package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/post"
	"github.com/labstack/echo/v4"
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

func (p PostHttpDelivery) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}
	err := p.interop.Delete(c.Request().Context(), token, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Delete success")

}

func (p PostHttpDelivery) GetDetail(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}
	post, err := p.interop.GetDetail(c.Request().Context(), token, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, post)

}
func (p PostHttpDelivery) GetByUid(c echo.Context) error {
	style := c.QueryParam("style")
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
	token := c.Request().Header.Get("Authorization")
	post, err := p.interop.GetByUid(c.Request().Context(), token, query, style)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (p PostHttpDelivery) GetOther(c echo.Context) error {
	uid := c.QueryParam("uid")
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
	token := c.Request().Header.Get("Authorization")
	post, err := p.interop.GetOther(c.Request().Context(), token, uid, query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (p PostHttpDelivery) Update(c echo.Context) error {
	postData := &post.Post{}
	if err := c.Bind(postData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := p.interop.Update(c.Request().Context(), token, postData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, postData)
}

func (p PostHttpDelivery) GetByCategory(c echo.Context) error {
	categoryId := c.QueryParam("category_id")
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
	token := c.Request().Header.Get("Authorization")
	post, err := p.interop.GetByCategory(c.Request().Context(), token, categoryId, query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (p PostHttpDelivery) UpdatePostComment(c echo.Context) error {
	id := c.QueryParam("id")
	postData := &post.Post{}
	if err := c.Bind(postData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := p.interop.UpdatePostComment(c.Request().Context(), token, id, postData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, postData)
}

func NewPostHttpDelivery(api *echo.Group, interop post.PostInterop) *PostHttpDelivery {
	handler := &PostHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	api.GET("/all", handler.List)
	api.GET("/detail", handler.GetDetail)
	api.GET("", handler.GetByUid)
	api.GET("/other", handler.GetOther)
	api.PUT("", handler.Update)
	api.GET("/category", handler.GetByCategory)
	api.DELETE("", handler.Delete)
	api.PUT("/comment", handler.UpdatePostComment)
	return handler
}
