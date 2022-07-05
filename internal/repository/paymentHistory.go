package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Final-Project-Kelompok-3/payments/internal/dto"
	"github.com/Final-Project-Kelompok-3/payments/internal/model"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type PaymentHistory interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.PaymentHistory, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.PaymentHistory, error)
	Create(ctx context.Context, paymentHistory model.PaymentHistory) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type paymentHistory struct {
	Db *gorm.DB
}

func NewPaymentHistory(db *gorm.DB) *paymentHistory {
	return &paymentHistory{db}
}

func (ph *paymentHistory) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.PaymentHistory, *dto.PaginationInfo, error) {
	
	var (
		paymentHistorys []model.PaymentHistory;
		count int64;
	)
	
	query := ph.Db.WithContext(ctx).Model(&model.PaymentHistory{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("LOWER(payment_code) LIKE ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&paymentHistorys).Error

	return paymentHistorys, dto.CheckInfoPagination(paginate,count), err
}

func (ph *paymentHistory) FindByID(ctx context.Context, ID uint) (model.PaymentHistory, error) {
	
	var paymentHistory model.PaymentHistory
	err := ph.Db.WithContext(ctx).Model(&paymentHistory).Where("id = ?", ID).First(&paymentHistory).Error

	return paymentHistory, err
}

func (ph *paymentHistory) Create(ctx context.Context, paymentHistory model.PaymentHistory) error {

	return ph.Db.WithContext(ctx).Model(&model.PaymentHistory{}).Create(&paymentHistory).Error
}

func (ph *paymentHistory) Update(ctx context.Context, ID uint, data map[string]interface{}) error {
	
	if data["password"] != nil {
		pss := fmt.Sprintf("%v", data["password"])
		bytes, _ := bcrypt.GenerateFromPassword([]byte(pss), bcrypt.DefaultCost)
		data["password"] = string(bytes)
	}

	err := ph.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.PaymentHistory{}).Updates(data).Error
	return err
}

func (ph *paymentHistory) Delete(ctx context.Context, ID uint) error {
	
	err := ph.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.PaymentHistory{}).Error
	return err
}