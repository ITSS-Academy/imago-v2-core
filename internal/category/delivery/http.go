package delivery

import (
	"github.com/itss-academy/imago/core/domain/category"
	"github.com/labstack/echo/v4"
)

type CategoryHttpDelivery struct {
	api     *echo.Group
	interop category.CategoryInterop
}

func (c CategoryHttpDelivery) Create(ctx echo.Context) error {
	categoryData := &category.Category{}
	if err := ctx.Bind(categoryData); err != nil {
		return ctx.NoContent(400)
	}
	err := c.interop.Create(ctx.Request().Context(), categoryData)
	if err != nil {
		return ctx.NoContent(500)
	}
	return ctx.JSON(201, categoryData)
}

func NewCategoryHttpDelivery(api *echo.Group, interop category.CategoryInterop) *CategoryHttpDelivery {
	handler := &CategoryHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	return handler
}
