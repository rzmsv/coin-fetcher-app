package middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ZapMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			logger.Info("HTTP request",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", res.Status),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
			)

			err := next(c)
			if err != nil {
				logger.Error("Handler error", zap.Error(err))
			}

			return err
		}
	}
}
