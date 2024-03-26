package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/report"
	"github.com/labstack/echo/v4"
)

type ReportHttpDelivery struct {
	api     *echo.Group
	interop report.ReportInterop
}

func (r ReportHttpDelivery) Create(c echo.Context) error {
	reportData := &report.Report{}
	if err := c.Bind(reportData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := r.interop.Create(c.Request().Context(), token, reportData)
	if err != nil {
		if errors.Is(err, report.ErrCreatorIDEmpty) || errors.Is(err, report.ErrReasonEmpty) || errors.Is(err, report.ErrContentEmpty) || errors.Is(err, report.ErrTypeEmpty) || errors.Is(err, report.ErrTypeIDEmpty) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, report.ErrReportNotCreated) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusCreated, reportData)
}

func (r ReportHttpDelivery) GetById(c echo.Context) error {
	id := c.QueryParam("id")
	fmt.Println(id)
	token := c.Request().Header.Get("Authorization")
	reportData, err := r.interop.GetById(c.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)

}

func (r ReportHttpDelivery) Get(c echo.Context) error {
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
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "page is invalid")
		}
		query.Page = page
	}
	if sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "size is invalid")
		}
		query.Size = size
	}
	reportData, err := r.interop.Get(c.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)

}

// GetAllByStatusCompleted
func (r ReportHttpDelivery) GetAllByStatusCompleted(c echo.Context) error {
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
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "page is invalid")
		}
		query.Page = page
	}
	if sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "size is invalid")
		}
		query.Size = size
	}
	reportData, err := r.interop.GetAllByStatusApproved(c.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)
}

// GetAllByStatusPending
func (r ReportHttpDelivery) GetAllByStatusPending(c echo.Context) error {
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
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "page is invalid")
		}
		query.Page = page
	}
	if sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "size is invalid")
		}
		query.Size = size
	}
	reportData, err := r.interop.GetAllByStatusPending(c.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)
}

func (r ReportHttpDelivery) Update(c echo.Context) error {
	id := c.QueryParam("id")
	reportData := &report.Report{}
	if err := c.Bind(reportData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := r.interop.Update(c.Request().Context(), token, reportData, id)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)
}

// ChangeStatusCompleted
func (r ReportHttpDelivery) ChangeStatusCompleted(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	err := r.interop.ChangeStatusApproved(c.Request().Context(), token, id, report.StatusApproved)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, true)
}

// ChangeStatusRejected
func (r ReportHttpDelivery) ChangeStatusRejected(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	err := r.interop.ChangeStatusRejected(c.Request().Context(), token, id, report.StatusRejected)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, true)

}

// Delete by id
func (r ReportHttpDelivery) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	err := r.interop.Delete(c.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, true)
}
func NewReportHttpDeliver(api *echo.Group, interop report.ReportInterop) *ReportHttpDelivery {
	handler := &ReportHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	api.GET("", handler.Get)
	api.GET("/id", handler.GetById)
	api.GET("/approved", handler.GetAllByStatusCompleted)
	api.GET("/pending", handler.GetAllByStatusPending)
	api.PUT("", handler.Update)
	api.PUT("/approved", handler.ChangeStatusCompleted)
	api.PUT("/rejected", handler.ChangeStatusRejected)
	api.DELETE("", handler.Delete)
	return handler
}
