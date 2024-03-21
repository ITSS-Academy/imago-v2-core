package delivery

import (
	"context"
	"github.com/itss-academy/imago/core/domain/profile"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileHttpDelivery struct {
	api     *echo.Group
	interop profile.ProfileInterop
}

func (p ProfileHttpDelivery) GetById(ctx context.Context, token string, id string) (*profile.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileHttpDelivery) GetAll(ctx context.Context, token string) ([]*profile.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileHttpDelivery) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}

	profileData := &profile.Profile{}
	if err := c.Bind(profileData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := p.interop.Create(c.Request().Context(), token, profileData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, profileData)
}

func (p ProfileHttpDelivery) Update(ctx context.Context, token string, profile *profile.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileHttpDelivery) Follow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileHttpDelivery) Unfollow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileHttpDelivery(api *echo.Group, interop profile.ProfileInterop) *ProfileHttpDelivery {
	handler := &ProfileHttpDelivery{
		api:     api,
		interop: interop,
	}
	api.POST("/mine", handler.Create)
	return handler
}
