package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentHistory struct {
	Model

	PaymentID     		int 	`json:"payment_id" gorm:"not null"`
	PaymentCode     	string 	`json:"payment_code" gorm:"size:100;not null"`
	PaymentAmount     	int 	`json:"payment_amount" gorm:"not null"`
	ExternalID    		string 	`json:"external_id" gorm:"size:200;not null"`
}

// BeforeCreate is a method for struct PaymentHistory
// gorm call this method before they execute query
func (model *PaymentHistory) BeforeCreate(tx *gorm.DB) (err error) {
	model.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct PaymentHistory
// gorm call this method before they execute query
func (model *PaymentHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	model.UpdatedAt = time.Now()
	return
}