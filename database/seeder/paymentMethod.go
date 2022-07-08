package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/payments/internal/model"
	"gorm.io/gorm"
)

func paymentMethodTableSeeder(conn *gorm.DB) {

	var users = []model.PaymentMethod{
		{Type: "Retail Outlets", MerchantCode: "ALFAMART", BankCode: "ALFAMART", BankName: "Alfamart"},
		{Type: "Retail Outlets", MerchantCode: "INDOMARET", BankCode: "INDOMARET", BankName: "Indomaret"},
	}

	if err := conn.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data payment method")
}