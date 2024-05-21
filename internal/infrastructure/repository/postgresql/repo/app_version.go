package repo

import (
	"context"

	"Booking/api-service-booking/internal/entity"
)

type AppVersionRepo interface {
	Get(ctx context.Context) (*entity.AppVersion, error)
	Create(ctx context.Context, m *entity.AppVersion) error
	Update(ctx context.Context, m *entity.AppVersion) error
}
