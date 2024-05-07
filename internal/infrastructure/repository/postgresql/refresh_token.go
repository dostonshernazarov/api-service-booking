package postgresql

import (
	"Booking/api-service-booking/internal/pkg/postgres"
	"Booking/api-service-booking/internal/usecase/refresh_token"
)

func NewRefreshTokenRepo(db *postgres.PostgresDB) refresh_token.RefreshTokenRepo {
	return nil
}
