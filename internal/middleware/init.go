package middleware

import (
	"net/http"
	"os"

	"github.com/Final-Project-Kelompok-3/payments/pkg/util/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
			Format:           "\n${time_custom} | ${status} | ${method} | ${host}${uri} | ${latency_human} | ${remote_ip} \n\n",
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)

	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}

}
