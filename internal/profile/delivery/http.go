package delivery

import (
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/profile"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type ProfileHttpDelivery struct {
	api     *echo.Group
	interop profile.ProfileInterop
}

func (p ProfileHttpDelivery) GetById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	//if token is empty return error
	if token == "" {
		return profile.ErrTokenEmpty
	}
	//using query param to get id
	profileData, err := p.interop.GetById(c.Request().Context(), token, c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) GetMine(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	//if token is empty return error
	if token == "" {
		return profile.ErrTokenEmpty
	}
	profileData, err := p.interop.GetMine(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	//if token is empty return error
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}
	profiles, err := p.interop.GetAll(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profiles)
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

func (p ProfileHttpDelivery) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}

	profileData := &profile.Profile{}
	if err := c.Bind(profileData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := p.interop.Update(c.Request().Context(), token, profileData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) Follow(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}

	profileId := c.QueryParam("profileId")
	profileOtherId := c.QueryParam("profileOtherId")

	err := p.interop.Follow(c.Request().Context(), token, profileId, profileOtherId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, "Followed")

}

func (p ProfileHttpDelivery) Unfollow(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}

	profileId := c.QueryParam("profileId")
	profileOtherId := c.QueryParam("profileOtherId")

	err := p.interop.Unfollow(c.Request().Context(), token, profileId, profileOtherId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, "Unfollowed")
}

func (p ProfileHttpDelivery) GetAllAuthProfile(e echo.Context) error {
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
	data, err := p.interop.GetAllAuthProfile(e.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, profile.ErrProfileNotFound) {
			return e.JSON(http.StatusNotFound, err.Error())
		}
		return e.JSON(http.StatusUnauthorized, err.Error())
	}
	return e.JSON(http.StatusOK, data)
}

func NewProfileHttpDelivery(api *echo.Group, interop profile.ProfileInterop) *ProfileHttpDelivery {
	handler := &ProfileHttpDelivery{
		api:     api,
		interop: interop,
	}
	api.GET("/all", handler.GetAll)
	api.GET("", handler.GetById)
	api.GET("/mine", handler.GetMine)
	api.GET("/authprofile", handler.GetAllAuthProfile)
	api.POST("/mine", handler.Create)
	api.PUT("/mine", handler.Update)
	api.PUT("/follow", handler.Follow)
	api.PUT("/unfollow", handler.Unfollow)
	return handler
}
