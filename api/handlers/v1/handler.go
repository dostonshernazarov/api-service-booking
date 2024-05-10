package v1

import (
	"time"

	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"

	grpcClients "Booking/api-service-booking/internal/infrastructure/grpc_service_client"
	"Booking/api-service-booking/internal/pkg/config"
	tokens "Booking/api-service-booking/internal/pkg/token"

	appV "Booking/api-service-booking/internal/usecase/app_version"
	"Booking/api-service-booking/internal/usecase/event"
	"Booking/api-service-booking/internal/usecase/refresh_token"
)

type HandlerV1 struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	jwtHandler     tokens.JwtHandler
	Service        grpcClients.ServiceClient
	RefreshToken   refresh_token.RefreshToken
	AppVersion     appV.AppVersion
	BrokerProducer event.BrokerProducer
	Enforcer       *casbin.Enforcer
}

type HandlerV1Config struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	RefreshToken   refresh_token.RefreshToken
	AppVersion     appV.AppVersion
	BrokerProducer event.BrokerProducer
	Enforcer       *casbin.Enforcer
}

func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
		Config:         c.Config,
		Logger:         c.Logger,
		ContextTimeout: c.ContextTimeout,
		Service:        c.Service,
		RefreshToken:   c.RefreshToken,
		AppVersion:     c.AppVersion,
		BrokerProducer: c.BrokerProducer,
		Enforcer:       c.Enforcer,
	}
}
