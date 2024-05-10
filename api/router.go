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

// @title welcome to Booking API
// @version 1.7
// @host localhost:8080

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	apiUser := api.Group("/users")


	// USER METHODS
	apiUser.POST("/create", HandlerV1.Create)
	apiUser.GET("/:id", HandlerV1.Get)
	apiUser.GET("/list/users", HandlerV1.ListUsers)
	apiUser.GET("/list/deleted", HandlerV1.ListDeletedUsers)
	apiUser.PUT("/update", HandlerV1.Update)
	apiUser.DELETE("/delete/:id", HandlerV1.Delete)

	// ATTRACTION METHODS
	api.POST("/attraction/create", HandlerV1.CreateAttraction)
	api.GET("/attraction/get", HandlerV1.GetAttraction)
	api.GET("/attraction/list", HandlerV1.ListAttractions)
	api.PATCH("/attraction/update", HandlerV1.UpdateAttraction)
	api.DELETE("/attraction/delete", HandlerV1.DeleteAttraction)

	// HOTEL METHODS
	api.POST("/hotel/create", HandlerV1.CreateHotel)
	api.GET("/hotel/get", HandlerV1.GetHotel)
	api.GET("/hotel/list", HandlerV1.ListHotels)
	api.PATCH("/hotel/update", HandlerV1.UpdateHotel)
	api.DELETE("/hotel/delete", HandlerV1.DeleteHotel)

	// REGITER METHODS
	api.POST("/users/register", HandlerV1.RegisterUser)
	api.GET("/users/verify", HandlerV1.Verification)
	api.GET("/users/login", HandlerV1.Login)
	api.GET("/users/set/:id", HandlerV1.ForgetPassword)
	api.GET("/users/code", HandlerV1.ForgetPasswordVerify)
	api.PUT("/users/password", HandlerV1.SetNewPassword)

	// ADMIN METHODS
	api.POST("/admins", HandlerV1.CreateAdmin)
	api.GET("/admins/:id", HandlerV1.GetAdmin)
	api.GET("/admins/list", HandlerV1.ListAdmins)
	api.PUT("/admins", HandlerV1.UpdateAdmin)
	api.DELETE("/admins/:id", HandlerV1.DeleteAdmin)



	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
