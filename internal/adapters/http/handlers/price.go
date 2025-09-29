package handler

import (
	"net/http"

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

func (a *Handler) GetPriceHandler(c echo.Context) error {
	symbol := c.Param("symbol")
	interval := c.Param("interval")
	if interval == "1min" {
		price, err := a.service.GetLastPrice(symbol)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No data"})
		}
		return c.JSON(http.StatusOK, map[string]float64{"price": price.Price})
	}
	price, err := a.service.GetAveragePrice(interval)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]float64{"price": price})
}
