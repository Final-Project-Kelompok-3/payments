package factory

import (
	"github.com/Final-Project-Kelompok-3/payments/internal/repository"
	"gorm.io/gorm"
)

type Factory struct {
	PaymentRepository repository.Payment
	PaymentHistoryRepository repository.PaymentHistory
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		PaymentRepository: repository.NewPayment(db),
		PaymentHistoryRepository: repository.NewPaymentHistory(db),
	}
}