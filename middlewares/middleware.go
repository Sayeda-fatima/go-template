package middlewares

import (
	"go-echo-template/common"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		common.Logger.LogInfo().Fields(map[string]interface {
		}{
			"method": c.Request().Method,
			"uri":    c.Request().URL.Path,
			"query":  c.Request().URL.RawQuery,
		}).Msg("Request")
		err := next(c)
		if err != nil {
			common.Logger.LogError().Fields(map[string]interface{}{
				"error": err.Error(),
			}).Msg("Response")
			return err
		}
		return nil
	}
}
