package rest

import (
	"developerProfile/internal/adapters/repository/postgres/developer_profile"
	"time"
)

type CreateProfileInDTO struct {
	Name                 string `json:"name"                   binding:"required"`
	Email                string `json:"email"                  binding:"required,email"`
	Phone                string `json:"phone"`
	Address              string `json:"address"`
	AcceptTermsOfService bool   `json:"accept_terms_of_service"`
}

type UpdateProfileInDTO struct {
	Name                 string `json:"name"                   binding:"required"`
	Email                string `json:"email"                  binding:"required,email"`
	Phone                string `json:"phone"`
	Address              string `json:"address"`
	AcceptTermsOfService bool   `json:"accept_terms_of_service"`
}

type ProfileOutDTO struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Phone                *string   `json:"phone,omitempty"`
	Address              *string   `json:"address,omitempty"`
	AcceptTermsOfService bool      `json:"accept_terms_of_service"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type ListProfilesOutDTO struct {
	Items []ProfileOutDTO `json:"items"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
}

type ErrorResponseDTO struct {
	Error string `json:"error"`
}

func toProfileOutDTO(p developer_profile.DeveloperProfile) ProfileOutDTO {
	dto := ProfileOutDTO{
		ID:                   p.ID.String(),
		Name:                 p.Name,
		Email:                p.Email,
		AcceptTermsOfService: p.AcceptTermsOfService,
		CreatedAt:            p.CreatedAt,
		UpdatedAt:            p.UpdatedAt,
	}
	if p.Phone.Valid {
		dto.Phone = &p.Phone.String
	}
	if p.Address.Valid {
		dto.Address = &p.Address.String
	}
	return dto
}