package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/Report"
	"github.com/labstack/echo/v4"
)

type ReportHttpDelivery struct {
	api     *echo.Group
	interop Report.ReportInterop
}

func (r ReportHttpDelivery) Create(c echo.Context) error {
	reportData := &Report.Report{}
	if err := c.Bind(reportData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := r.interop.Create(c.Request().Context(), token, reportData)
	if err != nil {
		if errors.Is(err, Report.ErrCreatorIDEmpty) || errors.Is(err, Report.ErrReasonEmpty) || errors.Is(err, Report.ErrContentEmpty) || errors.Is(err, Report.ErrTypeEmpty) || errors.Is(err, Report.ErrTypeIDEmpty) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if errors.Is(err, Report.ErrReportNotCreated) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusCreated, reportData)
}

func (r ReportHttpDelivery) GetById(c echo.Context) error {
	id := c.QueryParam("id")
	token := c.Request().Header.Get("Authorization")
	reportData, err := r.interop.GetById(c.Request().Context(), token, id)
	if err != nil {
		if errors.Is(err, Report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)

}

// get
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
		if errors.Is(err, Report.ErrReportNotFound) {
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
	reportData, err := r.interop.GetAllByStatusCompleted(c.Request().Context(), token, query)
	if err != nil {
		if errors.Is(err, Report.ErrReportNotFound) {
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
		if errors.Is(err, Report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)
}

func (r ReportHttpDelivery) Update(c echo.Context) error {
	reportData := &Report.Report{}
	if err := c.Bind(reportData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	token := c.Request().Header.Get("Authorization")
	err := r.interop.Update(c.Request().Context(), token, reportData)
	if err != nil {
		if errors.Is(err, Report.ErrReportNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, reportData)
}

func NewReportHttpDeliver(api *echo.Group, interop Report.ReportInterop) *ReportHttpDelivery {
	handler := &ReportHttpDelivery{api: api, interop: interop}
	api.POST("", handler.Create)
	api.GET("", handler.Get)
	api.GET("/completed", handler.GetAllByStatusCompleted)
	api.GET("/pending", handler.GetAllByStatusPending)
	api.GET("/id", handler.GetById)
	api.PUT("", handler.Update)
	return handler
}
