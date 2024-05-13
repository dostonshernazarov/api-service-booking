package api

import (
	// "net/http"
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
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	JwtHandler   tokens.JwtHandler
	BrokerProducer event.BrokerProducer
	AppVersion     app_version.AppVersion
	Enforcer       *casbin.Enforcer
}

// NewRouter
// @title Welcome To Booking API
// @Description API for Touristan
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRoute(option RouteOption) *gin.Engine {

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

	// REGISTER METHODS
	api.POST("/users/register", HandlerV1.RegisterUser)
	api.GET("/users/verify", HandlerV1.Verification)
	api.GET("/users/login", HandlerV1.Login)
	api.GET("/users/set/:id", HandlerV1.ForgetPassword)
	api.GET("/users/code", HandlerV1.ForgetPasswordVerify)
	api.PUT("/users/password", HandlerV1.SetNewPassword)

	api.GET("/token/:refresh", HandlerV1.UpdateToken)

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
