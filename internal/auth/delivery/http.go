package delivery

import (
	"errors"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/labstack/echo/v4"
	"net/http"
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

func NewAuthHttpDelivery(api *echo.Group, interop auth.AuthInterop) *AuthHttpDelivery {
	handler := &AuthHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	return handler
}
