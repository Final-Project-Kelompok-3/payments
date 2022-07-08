package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	Model

	Type     		string 	`json:"type" gorm:"size:100;not null"`
	MerchantCode     	string 	`json:"merchant_code" gorm:"size:100;not null"`
	BankCode     	string 	`json:"bank_code" gorm:"size:100;not null"`
	BankName     	string 	`json:"bank_name" gorm:"size:255;not null"`
}

// BeforeCreate is a method for struct PaymentMethod
// gorm call this method before they execute query
func (model *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {
	model.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct PaymentMethod
// gorm call this method before they execute query
func (model *PaymentMethod) BeforeUpdate(tx *gorm.DB) (err error) {
	model.UpdatedAt = time.Now()
	return
}