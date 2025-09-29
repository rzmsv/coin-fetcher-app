package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/username/coin-fetcher-app/internal/application"
)

type Handler struct {
	service *application.PriceService
	e       *echo.Echo
}

func NewHandler(e *echo.Echo, service *application.PriceService) *Handler {
	return &Handler{
		service,
		e,
	}
}

func (a *Handler) GetCoinPrice(c echo.Context) error {
	coin := c.Param("coin")
	interval := c.QueryParam("interval")
	if interval == "1min" || interval == "" {
		result, err := a.service.GetLastPrice(strings.ToLower(coin))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No data"})
		}
		return c.JSON(http.StatusOK, map[string]float64{"price": result.Price})
	}
	result, err := a.service.GetAveragePrice(interval, coin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]float64{"price": result})
}
