package dto

type CreatePaymentRequest struct {
	ParticipantID int `json:"participant_id" validate:"required"`
	Bill          int `json:"bill" validate:"required,min=1"`
	Payment       int `json:"payment"`
	Status        int `json:"status"`
}

type UpdatePaymentRequest struct {
	ParticipantID *int `json:"participant_id"`
	Bill          *int `json:"bill"`
	Payment       *int `json:"payment"`
	Status        *int `json:"status"`
}