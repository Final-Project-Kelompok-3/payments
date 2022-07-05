package response

import "github.com/Final-Project-Kelompok-3/payments/internal/dto"

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}
