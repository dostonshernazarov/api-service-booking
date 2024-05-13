package api

import (
	"net/http"
	"time"

	_ "Booking/api-service-booking/api/docs"
	v1 "Booking/api-service-booking/api/handlers/v1"

	// "Booking/api-service-booking/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/zap"

	grpcClients "Booking/api-service-booking/internal/infrastructure/grpc_service_client"
	"Booking/api-service-booking/internal/pkg/config"
	tokens "Booking/api-service-booking/internal/pkg/token"
	"Booking/api-service-booking/internal/usecase/app_version"
	"Booking/api-service-booking/internal/usecase/event"
	// "Booking/api-service-booking/internal/usecase/refresh_token"
)

type RouteOption struct {
	Config         	*config.Config
	Logger         	*zap.Logger
	ContextTimeout 	time.Duration
	Service        	grpcClients.ServiceClient
	JwtHandler   	tokens.JwtHandler
	BrokerProducer 	event.BrokerProducer
	AppVersion     	app_version.AppVersion
	Enforcer       	*casbin.Enforcer
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
		JwtHandler:   option.JwtHandler,
		AppVersion:     option.AppVersion,
		BrokerProducer: option.BrokerProducer,
		Enforcer:       option.Enforcer,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	// router.Use(middleware.Tracing)
	// router.Use(middleware.CheckCasbinPermission(option.Enforcer, *option.Config))

	router.Static("/media", "./media")

	api := router.Group("/v1")

	// BOOKING METHODS
		apiBooking := api.Group("/booking")
		apiBooking.POST("/create/hotel", HandlerV1.UHBCreate)
		apiBooking.POST("/create/restaurant", HandlerV1.URBCreate)
		apiBooking.POST("/create/attraction", HandlerV1.UABCreate)

		apiBooking.GET("/get/hotels/by/user/:id", HandlerV1.UHBGetAllByUId)
		apiBooking.GET("/get/restaurants/by/user/:id", HandlerV1.URBGetAllByUId)
		apiBooking.GET("/get/attractions/by/user/:id", HandlerV1.UABGetAllByUId)

		apiBooking.GET("/get/users/by/room/:id", HandlerV1.UHBGetAllByHId)
		apiBooking.GET("/get/users/by/restaurant/:id", HandlerV1.URBGetAllByRId)
		apiBooking.GET("/get/users/by/attraction/:id", HandlerV1.UABGetAllByAId)

		apiBooking.GET("/list/hotels", HandlerV1.UHBList)
		apiBooking.GET("/list/restaurants", HandlerV1.URBList)
		apiBooking.GET("/list/attractions", HandlerV1.UABList)

		apiBooking.GET("/list/deleted/hotels", HandlerV1.UHBListDeleted)
		apiBooking.GET("/list/deleted/restaurants", HandlerV1.URBListDeleted)
		apiBooking.GET("/list/deleted/attractions", HandlerV1.UABListDeleted)

		apiBooking.PUT("/update/booked/hotel", HandlerV1.UHBUpdate)
		apiBooking.PUT("/update/booked/restaurant", HandlerV1.URBUpdate)
		apiBooking.PUT("/update/booked/attraction", HandlerV1.UABUpdate)

		apiBooking.DELETE("/delete/hotel/:id", HandlerV1.UHBDelete)
		apiBooking.DELETE("/delete/restaurant/:id", HandlerV1.URBDelete)
		apiBooking.DELETE("/delete/attraction/:id", HandlerV1.UABDelete)
	
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
