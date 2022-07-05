package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	Model

	ParticipantID   int `json:"participan_id" gorm:"not null"`
	Bill    		int `json:"bill" gorm:"not null"`
	Payment    		int `json:"payment" gorm:"not null"`
	Status    		int `json:"status" gorm:"not null"`
}

// BeforeCreate is a method for struct Payment
// gorm call this method before they execute query
func (model *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	model.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Payment
// gorm call this method before they execute query
func (model *Payment) BeforeUpdate(tx *gorm.DB) (err error) {
	model.UpdatedAt = time.Now()
	return
}