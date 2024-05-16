package v1

import (
	"Booking/api-service-booking/internal/pkg/config"
	tokens "Booking/api-service-booking/internal/pkg/token"
	"net/http"
	"strings"

	"github.com/spf13/cast"
)

func GetIdFromToken(r *http.Request, cfg *config.Config) (string, int) {
	var softToken string
	token := r.Header.Get("Authorization")

	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		softToken = strings.TrimPrefix(token, "Bearer ")
	} else {
		softToken = token
	}

	claims, err := tokens.ExtractClaim(softToken, []byte(cfg.Token.SignInKey))
	if err != nil {
		return "unauthorized", http.StatusUnauthorized
	}

	return cast.ToString(claims["sub"]), 200
}
