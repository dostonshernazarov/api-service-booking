package v1

import (
	"time"

	"go.uber.org/zap"

	grpcClients "Booking/api_establishment_booking/internal/infrastructure/grpc_service_client"
	"Booking/api_establishment_booking/internal/pkg/config"

	appV "Booking/api_establishment_booking/internal/usecase/app_version"
	"Booking/api_establishment_booking/internal/usecase/event"
	"Booking/api_establishment_booking/internal/usecase/refresh_token"
)

type HandlerV1 struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	RefreshToken   refresh_token.RefreshToken
	AppVersion     appV.AppVersion
	BrokerProducer event.BrokerProducer
}

type HandlerV1Config struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	RefreshToken   refresh_token.RefreshToken
	AppVersion     appV.AppVersion
	BrokerProducer event.BrokerProducer
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
	}
}
