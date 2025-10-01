package http

import (
	"github.com/labstack/echo/v4"
	handler "github.com/username/coin-fetcher-app/internal/adapters/http/handlers"
	"github.com/username/coin-fetcher-app/internal/adapters/http/routes"
	"github.com/username/coin-fetcher-app/internal/application"
	"github.com/username/coin-fetcher-app/internal/middleware"
	"github.com/username/coin-fetcher-app/pkg/logger"
)

type EchoAdapter struct {
	service *application.PriceService
	e       *echo.Echo
}

func NewEchoAdapter(service *application.PriceService) *EchoAdapter {
	e := echo.New()
	adapter := &EchoAdapter{service, e}
	price_handler := handler.NewHandler(e, service)
	routes.NewPriceRoute(price_handler).HandlerPriceRoutes(e)
	e.Use(middleware.ZapMiddleware(logger.LoggerZap()))
	return adapter
}

func (a *EchoAdapter) Start(addr string) error {
	return a.e.Start(addr)
}
