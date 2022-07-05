package payment

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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Payment], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Payment, error)
	Create(ctx context.Context, payload *dto.CreatePaymentRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdatePaymentRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Payment, error)
}

type service struct {
	PaymentRepository repository.Payment
}

func NewService(f *factory.Factory) Service {
	return &service{
		PaymentRepository: f.PaymentRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Payment], error) {
	
	Books, info, err := s.PaymentRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	
	result := new(dto.SearchGetResponse[model.Payment])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Payment, error) {

	data, err := s.PaymentRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreatePaymentRequest) (string, error) {
	
	var payment = model.Payment{
		ParticipantID: 	payload.ParticipantID,
		Bill:       	payload.Bill,
		Payment: 		payload.Payment,
		Status: 		payload.Status,
	}

	err := s.PaymentRepository.Create(ctx, payment)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdatePaymentRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.ParticipantID != nil {
		data["participant_id"] = payload.ParticipantID
	}
	if payload.Bill != nil {
		data["bill"] = payload.Bill
	}
	if payload.Payment != nil {
		data["payment"] = payload.Payment
	}
	if payload.Status != nil {
		data["status"] = payload.Status
	}

	err := s.PaymentRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Payment, error) {
	data, err := s.PaymentRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.PaymentRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}