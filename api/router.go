package api

import (
	// "net/http"
	"time"

	_ "Booking/api-service-booking/api/docs"
	v1 "Booking/api-service-booking/api/handlers/v1"

	"Booking/api-service-booking/api/middleware"

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
// @securityDefinitions.apikey BearerAuth
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
	router.Use(middleware.CheckCasbinPermission(option.Enforcer, *option.Config))

	router.Static("/media", "./media")
	api := router.Group("/v1")

	// USER METHODS

	api.POST("/users", HandlerV1.Create)
	api.GET("/users/:id", HandlerV1.Get)
	api.GET("/users/list", HandlerV1.ListUsers)
	api.GET("/users/list/deleted", HandlerV1.ListDeletedUsers)
	api.PUT("/users", HandlerV1.Update)
	api.DELETE("/users/:id", HandlerV1.Delete)
	api.GET("/users/token", HandlerV1.GetByToken)

	// ATTRACTION METHODS
	api.POST("/attraction", HandlerV1.CreateAttraction)
	api.GET("/attraction", HandlerV1.GetAttraction)
	api.GET("/attraction/list", HandlerV1.ListAttractions)
	api.PUT("/attraction", HandlerV1.UpdateAttraction)
	api.DELETE("/attraction", HandlerV1.DeleteAttraction)
	api.GET("/attraction/listlocation", HandlerV1.ListAttractionsByLocation)
	api.GET("/attraction/find", HandlerV1.FindAttractionsByName)

	// HOTEL METHODS
	api.POST("/hotel", HandlerV1.CreateHotel)
	api.GET("/hotel", HandlerV1.GetHotel)
	api.GET("/hotel/list", HandlerV1.ListHotels)
	api.PUT("/hotel", HandlerV1.UpdateHotel)
	api.DELETE("/hotel", HandlerV1.DeleteHotel)
	api.GET("/hotel/listlocation", HandlerV1.ListHotelsByLocation)
	api.GET("/hotel/find", HandlerV1.FindHotelsByName)

	// RESTAURANT METHODS
	api.POST("/restaurant", HandlerV1.CreateRestaurant)
	api.GET("/restaurant", HandlerV1.GetRestaurant)
	api.GET("/restaurant/list", HandlerV1.ListRestaurants)
	api.PUT("/restaurant", HandlerV1.UpdateRestaurant)
	api.DELETE("/restaurant", HandlerV1.DeleteRestaurant)
	api.GET("/restaurant/listlocation", HandlerV1.ListRestaurantsByLocation)
	api.GET("/restaurant/find", HandlerV1.FindRestaurantsByName)

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
	api.POST("/users/login", HandlerV1.Login)
	api.GET("/users/set/:email", HandlerV1.ForgetPassword)
	api.GET("/users/code", HandlerV1.ForgetPasswordVerify)
	api.PUT("/users/password", HandlerV1.SetNewPassword)
	api.POST("/admins/login", HandlerV1.LoginAdmin)

	api.GET("/token/:refresh", HandlerV1.UpdateToken)

	// ADMIN METHODS
	api.POST("/admins", HandlerV1.CreateAdmin)
	api.GET("/admins/:id", HandlerV1.GetAdmin)
	api.GET("/admins/list", HandlerV1.ListAdmins)
	api.PUT("/admins", HandlerV1.UpdateAdmin)
	api.DELETE("/admins/:id", HandlerV1.DeleteAdmin)

	// MEDIA
	api.POST("/media/user-photo", HandlerV1.UploadMedia)
	api.POST("/media/establishment/:id", HandlerV1.CreateEstablishmentMedia)

	// BOOKING HOTEL
	api.POST("/booking/hotels", HandlerV1.UHBCreate)
	api.GET("/booking/hotels/:id", HandlerV1.UHBGetAllByUId)
	api.GET("/booking/users/room/:id", HandlerV1.UHBGetAllByHId)
	api.GET("/booking/hotels", HandlerV1.UHBList)
	api.GET("/booking/hotels/deleted", HandlerV1.UHBListDeleted)
	api.PUT("/booking/hotels", HandlerV1.UHBUpdate)
	api.DELETE("/booking/hotels/:id", HandlerV1.UHBDelete)

	// BOOKING RESTAURANT
	api.POST("/booking/restaurants", HandlerV1.URBCreate)
	api.GET("/booking/restaurants/:id", HandlerV1.URBGetAllByUId)
	api.GET("/booking/users/restaurant/:id", HandlerV1.URBGetAllByRId)
	api.GET("/booking/restaurants", HandlerV1.URBList)
	api.GET("/booking/restaurants/deleted", HandlerV1.URBListDeleted)
	api.PUT("/booking/restaurants", HandlerV1.URBUpdate)
	api.DELETE("/booking/restaurants/:id", HandlerV1.URBDelete)

	// BOOKING ATTRACTION
	api.POST("/booking/attractions", HandlerV1.UABCreate)
	api.GET("/booking/attractions/:id", HandlerV1.UABGetAllByUId)
	api.GET("/booking/users/attraction/:id", HandlerV1.UABGetAllByAId)
	api.GET("/booking/attractions", HandlerV1.UABList)
	api.GET("/booking/attractions/deleted", HandlerV1.UABListDeleted)
	api.PUT("/booking/attractions", HandlerV1.UABUpdate)
	api.DELETE("/booking/attractions/:id", HandlerV1.UABDelete)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
