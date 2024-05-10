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
	api.GET("/attraction/get", HandlerV1.GetAttraction)
	api.GET("/attraction/list", HandlerV1.ListAttractions)
	api.PATCH("/attraction/update", HandlerV1.UpdateAttraction)
	api.DELETE("/attraction/delete", HandlerV1.DeleteAttraction)
	api.GET("/attraction/listbylocation", HandlerV1.ListAttractionsByLocation)

	// HOTEL METHODS
	api.POST("/hotel/create", HandlerV1.CreateHotel)
	api.GET("/hotel/get", HandlerV1.GetHotel)
	api.GET("/hotel/list", HandlerV1.ListHotels)
	api.PATCH("/hotel/update", HandlerV1.UpdateHotel)
	api.DELETE("/hotel/delete", HandlerV1.DeleteHotel)

	// RESTAURANT METHODS
	api.POST("/restaurant/create", HandlerV1.CreateRestaurant)
	api.GET("/restaurant/get", HandlerV1.GetRestaurant)
	api.GET("/restaurant/list", HandlerV1.ListRestaurants)
	api.PATCH("/restaurant/update", HandlerV1.UpdateRestaurant)
	api.DELETE("/restaurant/delete", HandlerV1.DeleteRestaurant)

	// FAVOURITE METHODS
	api.POST("/favourite/add", HandlerV1.AddToFavourites)
	api.DELETE("/favourite/remove", HandlerV1.RemoveFromFavourites)
	api.GET("/favourite/list", HandlerV1.ListFavouritesByUserId)

	// REVIEW METHODS
	api.POST("/review/create", HandlerV1.CreateReview)
	api.GET("/review/list", HandlerV1.ListReviews)
	api.DELETE("/review/delete", HandlerV1.DeleteReview)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
