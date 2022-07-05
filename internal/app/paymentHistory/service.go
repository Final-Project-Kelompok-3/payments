package paymentHistory

import (
	"context"

	"github.com/Final-Project-Kelompok-3/payments/pkg/constant"
	res "github.com/Final-Project-Kelompok-3/payments/pkg/util/response"

	"github.com/Final-Project-Kelompok-3/payments/internal/dto"
	"github.com/Final-Project-Kelompok-3/payments/internal/factory"
	"github.com/Final-Project-Kelompok-3/payments/internal/model"
	"github.com/Final-Project-Kelompok-3/payments/internal/repository"
)

type Service interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.PaymentHistory], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.PaymentHistory, error)
	Create(ctx context.Context, payload *dto.CreatePaymentHistoryRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdatePaymentHistoryRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.PaymentHistory, error)
}

type service struct {
	PaymentHistoryRepository repository.PaymentHistory
}

func NewService(f *factory.Factory) Service {
	return &service{
		PaymentHistoryRepository: f.PaymentHistoryRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.PaymentHistory], error) {
	
	Books, info, err := s.PaymentHistoryRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	
	result := new(dto.SearchGetResponse[model.PaymentHistory])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.PaymentHistory, error) {

	data, err := s.PaymentHistoryRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreatePaymentHistoryRequest) (string, error) {
	
	var paymentHistory = model.PaymentHistory{
		PaymentID: payload.PaymentID,
		PaymentCode: payload.PaymentCode,
		PaymentAmount: payload.PaymentAmount,
		ExternalID: payload.ExternalID,
	}

	err := s.PaymentHistoryRepository.Create(ctx, paymentHistory)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdatePaymentHistoryRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.PaymentID != nil {
		data["payment_id"] = payload.PaymentID
	}
	if payload.PaymentCode != nil {
		data["payment_code"] = payload.PaymentCode
	}
	if payload.PaymentAmount != nil {
		data["payment_amount"] = payload.PaymentAmount
	}
	if payload.ExternalID != nil {
		data["external_id"] = payload.ExternalID
	}

	err := s.PaymentHistoryRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.PaymentHistory, error) {
	data, err := s.PaymentHistoryRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.PaymentHistoryRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}