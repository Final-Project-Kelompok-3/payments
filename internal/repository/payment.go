package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/payments/internal/dto"
	"github.com/Final-Project-Kelompok-3/payments/internal/model"

	"gorm.io/gorm"
)

type Payment interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Payment, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Payment, error)
	Create(ctx context.Context, payment model.Payment) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type payment struct {
	Db *gorm.DB
}

func NewPayment(db *gorm.DB) *payment {
	return &payment{db}
}

func (p *payment) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Payment, *dto.PaginationInfo, error) {
	
	var (
		payments []model.Payment;
		count int64;
	)
	
	query := p.Db.WithContext(ctx).Model(&model.Payment{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(status) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&payments).Error

	return payments, dto.CheckInfoPagination(paginate,count), err
}

func (p *payment) FindByID(ctx context.Context, ID uint) (model.Payment, error) {
	
	var payment model.Payment
	err := p.Db.WithContext(ctx).Model(&payment).Where("id = ?", ID).First(&payment).Error

	return payment, err
}

func (p *payment) Create(ctx context.Context, payment model.Payment) error {

	return p.Db.WithContext(ctx).Model(&model.Payment{}).Create(&payment).Error
}

func (p *payment) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Payment{}).Updates(data).Error
	return err
}

 func (p *payment) Delete(ctx context.Context, ID uint) error {
	err := p.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Payment{}).Error
	return err
 }