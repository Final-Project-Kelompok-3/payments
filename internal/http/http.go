package http

import (
	"github.com/Final-Project-Kelompok-3/payments/internal/app/payment"
	"github.com/Final-Project-Kelompok-3/payments/internal/app/paymentHistory"
	"github.com/Final-Project-Kelompok-3/payments/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1 := e.Group("payment/v1")
	payment.NewHandler(f).Route(v1.Group("/payments"))
	paymentHistory.NewHandler(f).Route(v1.Group("/paymentHistories"))
}
