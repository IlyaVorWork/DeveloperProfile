package service

import (
	"context"
	"database/sql"
	"developerProfile/internal/adapters/repository/postgres"
	"developerProfile/internal/adapters/repository/postgres/developer_profile"

	"github.com/google/uuid"
)

type CreateProfileParams struct {
	Name                 string
	Email                string
	Phone                string
	Address              string
	AcceptTermsOfService bool
}

type UpdateProfileParams struct {
	Name                 string
	Email                string
	Phone                string
	Address              string
	AcceptTermsOfService bool
}

type Service struct {
	repo postgres.DeveloperProfileRepository
}

func NewService(repo postgres.DeveloperProfileRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProfile(ctx context.Context, p CreateProfileParams) (developer_profile.DeveloperProfile, error) {
	return s.repo.CreateDeveloperProfile(ctx, developer_profile.CreateDeveloperProfileParams{
		Name:                 p.Name,
		Email:                p.Email,
		Phone:                sql.NullString{String: p.Phone, Valid: p.Phone != ""},
		Address:              sql.NullString{String: p.Address, Valid: p.Address != ""},
		AcceptTermsOfService: p.AcceptTermsOfService,
	})
}

func (s *Service) GetProfile(ctx context.Context, id string) (developer_profile.DeveloperProfile, error) {
	return s.repo.GetDeveloperProfile(ctx, uuid.MustParse(id))
}

func (s *Service) ListProfiles(ctx context.Context, page, size int) ([]developer_profile.DeveloperProfile, error) {
	list, err := s.repo.ListDeveloperProfiles(ctx, developer_profile.ListDeveloperProfilesParams{
		Limit:  int32(size),
		Offset: int32(page * size),
	})
	if err != nil {
		return nil, err
	}
	if list == nil {
		return []developer_profile.DeveloperProfile{}, nil
	}
	return list, nil
}

func (s *Service) UpdateProfile(ctx context.Context, id string, p UpdateProfileParams) (developer_profile.DeveloperProfile, error) {
	return s.repo.UpdateDeveloperProfile(ctx, developer_profile.UpdateDeveloperProfileParams{
		ID:                   uuid.MustParse(id),
		Name:                 p.Name,
		Email:                p.Email,
		Phone:                sql.NullString{String: p.Phone, Valid: p.Phone != ""},
		Address:              sql.NullString{String: p.Address, Valid: p.Address != ""},
		AcceptTermsOfService: p.AcceptTermsOfService,
	})
}

func (s *Service) DeleteProfile(ctx context.Context, id string) error {
	return s.repo.DeleteDeveloperProfile(ctx, uuid.MustParse(id))
}