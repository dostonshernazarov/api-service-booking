package api

import (
	"net/http"
	"time"

	_ "Booking/api-service-booking/api/docs"
	v1 "Booking/api-service-booking/api/handlers/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/zap"

	grpcClients "Booking/api-service-booking/internal/infrastructure/grpc_service_client"
	"Booking/api-service-booking/internal/pkg/config"
	"Booking/api-service-booking/internal/usecase/app_version"
	"Booking/api-service-booking/internal/usecase/event"
	"Booking/api-service-booking/internal/usecase/refresh_token"
)

type RouteOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	RefreshToken   refresh_token.RefreshToken
	BrokerProducer event.BrokerProducer
	AppVersion     app_version.AppVersion
}

func NewRoute(option RouteOption) http.Handler {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Service:        option.Service,
		RefreshToken:   option.RefreshToken,
		AppVersion:     option.AppVersion,
		BrokerProducer: option.BrokerProducer,
	})

	api := router.Group("/v1")

	// ATTRACTION METHODS
	api.POST("/attraction/create", HandlerV1.CreateAttraction)

	// REGITER METHODS
	api.POST("/users/register", HandlerV1.RegisterUser)
	api.GET("/users/verify", HandlerV1.Verification)
	api.GET("/users/login", HandlerV1.Login)
	api.GET("/users/set/:id", HandlerV1.ForgetPassword)
	api.GET("/users/code", HandlerV1.ForgetPasswordVerify)
	api.PUT("/users/password", HandlerV1.SetNewPassword)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
