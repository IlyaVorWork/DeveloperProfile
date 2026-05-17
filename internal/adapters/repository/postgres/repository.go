package postgres

import (
	"context"
	"developerProfile/internal/adapters/repository/postgres/developer_profile"

	"github.com/google/uuid"
)

type DeveloperProfileRepository interface {
	CreateDeveloperProfile(ctx context.Context, arg developer_profile.CreateDeveloperProfileParams) (developer_profile.DeveloperProfile, error)
	GetDeveloperProfile(ctx context.Context, id uuid.UUID) (developer_profile.DeveloperProfile, error)
	ListDeveloperProfiles(ctx context.Context, arg developer_profile.ListDeveloperProfilesParams) ([]developer_profile.DeveloperProfile, error)
	UpdateDeveloperProfile(ctx context.Context, arg developer_profile.UpdateDeveloperProfileParams) (developer_profile.DeveloperProfile, error)
	DeleteDeveloperProfile(ctx context.Context, id uuid.UUID) error
}
