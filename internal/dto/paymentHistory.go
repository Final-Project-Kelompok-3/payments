package dto

type CreatePaymentHistoryRequest struct {
	PaymentID     int    `json:"payment_id" validate:"required"`
	PaymentCode   string `json:"payment_code" validate:"required"`
	PaymentAmount int    `json:"payment_amount" validate:"required"`
	ExternalID    string `json:"external_id"  validate:"required"`
}

type UpdatePaymentHistoryRequest struct {
	PaymentID     *int    `json:"payment_id"`
	PaymentCode   *string `json:"payment_code"`
	PaymentAmount *int    `json:"payment_amount"`
	ExternalID    *string `json:"external_id"`
}

type PaymentHistoryResponse struct {
	ID            int    `json:"id"`
	PaymentID     int    `json:"payment_id"`
	PaymentCode   string `json:"payment_code"`
	PaymentAmount int    `json:"payment_amount"`
	ExternalID    string `json:"external_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"deleted_at"`
}