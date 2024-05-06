package postgresql

import (
	"Booking/api_establishment_booking/internal/pkg/postgres"
	"Booking/api_establishment_booking/internal/usecase/refresh_token"
)

func NewRefreshTokenRepo(db *postgres.PostgresDB) refresh_token.RefreshTokenRepo {
	return nil
}
