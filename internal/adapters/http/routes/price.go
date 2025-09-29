package routes

import (
	"github.com/labstack/echo/v4"
	handler "github.com/username/coin-fetcher-app/internal/adapters/http/handlers"
)

type price_routes struct {
	priceHandler *handler.Handler
}

func NewPriceRoute(priceHandler *handler.Handler) *price_routes {
	return &price_routes{
		priceHandler,
	}
}

func (r *price_routes) HandlerPriceRoutes(e *echo.Echo) {

	api := e.Group("/api")
	api.GET("/price/history/:coin", r.priceHandler.GetCoinPrice)
}
