package delivery

import (
	"errors"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AuthHttpDelivery struct {
	api     *echo.Group
	interop auth.AuthInterop
}

func (a AuthHttpDelivery) Create(c echo.Context) error {
	authData := &auth.Auth{}
	if err := c.Bind(authData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := a.interop.Create(c.Request().Context(), token, authData)
	if err != nil {
		if errors.Is(err, auth.ErrEmailEmpty) || errors.Is(err, auth.ErrRoleIdEmpty) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, auth.ErrAuthNotCreated) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusCreated, authData)
}

func (a AuthHttpDelivery) GetById(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	authData, err := a.interop.GetById(c.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, auth.ErrAuthNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, authData)
}

func (a AuthHttpDelivery) Get(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
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
	authData, err := a.interop.Get(c.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, auth.ErrAuthNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())

	}
	return c.JSON(http.StatusOK, authData)
}

func (a AuthHttpDelivery) Update(c echo.Context) error {
	authData := &auth.Auth{}
	if err := c.Bind(authData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := a.interop.Update(c.Request().Context(), token, authData)
	if err != nil {
		if errors.Is(err, auth.ErrEmailEmpty) || errors.Is(err, auth.ErrRoleIdEmpty) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, auth.ErrAuthNotUpdated) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, authData)
}

func (a AuthHttpDelivery) ChangeRole(c echo.Context) error {
	roleId := c.QueryParam("role_id")
	token := c.Request().Header.Get("Authorization")
	err := a.interop.ChangeRole(c.Request().Context(), token, roleId)
	if err != nil {
		if errors.Is(err, auth.ErrAuthNotAuthorized) {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (a AuthHttpDelivery) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	err := a.interop.Delete(c.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, auth.ErrAuthNotDeleted) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func NewAuthHttpDelivery(api *echo.Group, interop auth.AuthInterop) *AuthHttpDelivery {
	handler := &AuthHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	api.GET("", handler.Get)
	api.GET("/id", handler.GetById)
	api.PUT("", handler.Update)
	api.DELETE("", handler.Delete)
	return handler
}
