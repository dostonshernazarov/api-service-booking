package v1

import (
	models "Booking/api-service-booking/api/models"
	pbb "Booking/api-service-booking/genproto/booking-proto"
	l "Booking/api-service-booking/internal/pkg/logger"
	"Booking/api-service-booking/internal/pkg/otlp"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Hotel Booking
// @Summary Create Hotel Booking
// @Security BearerAuth
// @Description Api for Create Hotel Booking
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param CreateBookingReq body models.CreateBookingReq true "createModel"
// @Success 200 {object} models.BookingRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/hotels [post]
func (h *HandlerV1) UHBCreate(c *gin.Context) {

	ctx, span := otlp.Start(c, "api", "CreateBookHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.CreateBookingReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}


	userID, statusCode := GetIdFromToken(c.Request, h.Config)
    if statusCode == 401 {
		c.JSON(http.StatusUnauthorized, models.Error{
			Message: "Log In Again",
		})
		return
	}

	response, err := h.Service.BookingService().UHBCreate(ctx, &pbb.GeneralBook{
		Id:                   uuid.NewString(),
		UserId:               userID,
		HraId:                body.HraId,
		WillArrive:           body.WillArrive,
		WillLeave:            body.WillLeave,
		NumberOfPeople:       body.NumberOfPeople,
		IsCanceled:           body.IsCanceled,
		Reason:               body.Reason,
		CreatedAt:            time.Now().Format("2006-01-02T15:04:05"),
	})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	c.JSON(http.StatusCreated, &models.BookingRes{
		Id:                   uuid.MustParse(response.Id),
		UserId:               response.UserId,
		HraId:                response.HraId,
		WillArrive:           response.WillArrive,
		WillLeave:            response.WillLeave,
		NumberOfPeople:       response.NumberOfPeople,
		IsCanceled:           response.IsCanceled,
		Reason:               response.Reason,
		CreatedAt:            response.CreatedAt,
	})
}

// Create Restaurant Booking
// @Summary Create Restaurant Booking
// @Security BearerAuth
// @Description Api for Create Restaurant Booking
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param CreateBookingReq body models.CreateBookingReq true "createModel"
// @Success 200 {object} models.BookingRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/restaurants [post]
func (h *HandlerV1) URBCreate(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "CreateBooRestaurant")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.CreateBookingReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	userID, statusCode := GetIdFromToken(c.Request, h.Config)
    if statusCode == 401 {
		c.JSON(http.StatusUnauthorized, models.Error{
			Message: "Log In Again",
		})
		return
	}

	response, err := h.Service.BookingService().URBCreate(ctx, &pbb.GeneralBook{
		Id:                   uuid.NewString(),
		UserId:               userID,
		HraId:                body.HraId,
		WillArrive:           body.WillArrive,
		WillLeave:            body.WillLeave,
		NumberOfPeople:       body.NumberOfPeople,
		IsCanceled:           body.IsCanceled,
		Reason:               body.Reason,
		CreatedAt:            time.Now().Format("2006-01-02T15:04:05"),
	})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	c.JSON(http.StatusCreated, &models.BookingRes{
		Id:                   uuid.MustParse(response.Id),
		UserId:               response.UserId,
		HraId:                response.HraId,
		WillArrive:           response.WillArrive,
		WillLeave:            response.WillLeave,
		NumberOfPeople:       response.NumberOfPeople,
		IsCanceled:           response.IsCanceled,
		Reason:               response.Reason,
		CreatedAt:            response.CreatedAt,
	})
}

// Create Attraction Booking
// @Summary Create Attraction Booking
// @Security BearerAuth
// @Description Api for Create Attraction Booking
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param CreateBookingReq body models.CreateBookingReq true "createModel"
// @Success 200 {object} models.BookingRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/attractions [post]
func (h *HandlerV1) UABCreate(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "CreateBookAttraction")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.CreateBookingReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	userID, statusCode := GetIdFromToken(c.Request, h.Config)
    if statusCode == 401 {
		c.JSON(http.StatusUnauthorized, models.Error{
			Message: "Log In Again",
		})
		return
	}

	response, err := h.Service.BookingService().UABCreate(ctx, &pbb.GeneralBook{
		Id:                   uuid.NewString(),
		UserId:               userID,
		HraId:                body.HraId,
		WillArrive:           body.WillArrive,
		WillLeave:            body.WillLeave,
		NumberOfPeople:       body.NumberOfPeople,
		IsCanceled:           body.IsCanceled,
		Reason:               body.Reason,
		CreatedAt:            time.Now().Format("2006-01-02T15:04:05"),
	})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	c.JSON(http.StatusCreated, &models.BookingRes{
		Id:                   uuid.MustParse(response.Id),
		UserId:               response.UserId,
		HraId:                response.HraId,
		WillArrive:           response.WillArrive,
		WillLeave:            response.WillLeave,
		NumberOfPeople:       response.NumberOfPeople,
		IsCanceled:           response.IsCanceled,
		Reason:               response.Reason,
		CreatedAt:            response.CreatedAt,
	})
}

// Get All Hotels By User Id
// @Summary Get All Hotels By User Id
// @Security BearerAuth
// @Description Api for Get All Hotels By User Id
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAllByUIdRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/hotels/{id} [get]
func (h *HandlerV1) UHBGetAllByUId(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookHotelByUserID")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.GetAllByUIdReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	id := c.Param("id")

	response, err := h.Service.BookingService().UHBGetAllByUId(
		ctx, &pbb.ListReqById{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
			Id:     &pbb.Id{
				Id: id,
			},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Get All Restaurants By User Id
// @Summary Get All Restaurants By User Id
// @Security BearerAuth
// @Description Api for Get All Restaurants By User Id
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAllByUIdRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/restaurants/{id} [get]
func (h *HandlerV1) URBGetAllByUId(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookRestaurantsByUserID")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.GetAllByUIdReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	id := c.Param("id")

	response, err := h.Service.BookingService().URBGetAllByUId(
		ctx, &pbb.ListReqById{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
			Id:     &pbb.Id{
				Id: id,
			},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Get All Attractions By User Id
// @Summary Get All Attractions By User Id
// @Security BearerAuth
// @Description Api for Get All Attractions By User Id
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAllByUIdRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/attractions/{id} [get]
func (h *HandlerV1) UABGetAllByUId(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookAttractionByUserID")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.GetAllByUIdReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	id := c.Param("id")

	response, err := h.Service.BookingService().UABGetAllByUId(
		ctx, &pbb.ListReqById{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
			Id:     &pbb.Id{
				Id: id,
			},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Get All Users By Room Id
// @Summary Get All Users By Room Id
// @Security BearerAuth
// @Description Api for Get All Users By Room Id
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAllByHRAIdRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/users/room/{id} [get]
func (h *HandlerV1) UHBGetAllByHId(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listUsersByBookedRoomID")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.GetAllByHRAIdReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	id := c.Param("id")

	response, err := h.Service.BookingService().UHBGetAllByHId(
		ctx, &pbb.ListReqById{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
			Id:     &pbb.Id{
				Id: id,
			},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Get All Users By Restaurant Id
// @Summary Get All Users By Restaurant Id
// @Security BearerAuth
// @Description Api for Get All Users By Restaurant Id
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAllByHRAIdRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/users/restaurant/{id} [get]
func (h *HandlerV1) URBGetAllByRId(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listUsersByRestaurantID")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.GetAllByHRAIdReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	id := c.Param("id")

	response, err := h.Service.BookingService().UHBGetAllByHId(
		ctx, &pbb.ListReqById{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
			Id:     &pbb.Id{
				Id: id,
			},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Get All Users By Attraction Id
// @Summary Get All Users By Attraction Id
// @Security BearerAuth
// @Description Api for Get All Users By Attraction Id
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAllByHRAIdRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/users/attraction/{id} [get]
func (h *HandlerV1) UABGetAllByAId(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookHotelByUserID")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.GetAllByHRAIdReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	id := c.Param("id")

	response, err := h.Service.BookingService().UABGetAllByAId(
		ctx, &pbb.ListReqById{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
			Id:     &pbb.Id{
				Id: id,
			},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// List Hotels
// @Summary List Hotels
// @Security BearerAuth
// @Description Api for List Hotels
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/hotels [get]
func (h *HandlerV1) UHBList(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookHotels")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.Pagination
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.BookingService().UHBList(
		ctx, &pbb.ListReq{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// List Restaurants
// @Summary List Restaurants
// @Security BearerAuth
// @Description Api for List Restaurants
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/restaurants [get]
func (h *HandlerV1) URBList(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookRestaurants")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.Pagination
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	response, err := h.Service.BookingService().URBList(
		ctx, &pbb.ListReq{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// List Attractions
// @Summary List Attractions
// @Security BearerAuth
// @Description Api for List Attractions
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/attractions [get]
func (h *HandlerV1) UABList(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listBookAttractions")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.Pagination
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	response, err := h.Service.BookingService().UABList(
		ctx, &pbb.ListReq{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// List Deleted Hotels
// @Summary List Deleted Hotels
// @Security BearerAuth
// @Description Api for List Deleted Hotels
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/hotels/deleted [get]
func (h *HandlerV1) UHBListDeleted(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listDeletedBookHotels")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.Pagination
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	response, err := h.Service.BookingService().UHBListDeleted(
		ctx, &pbb.ListReq{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// List Deleted Restaurants
// @Summary List Deleted Restaurants
// @Security BearerAuth
// @Description Api for List Deleted Restaurants
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param irequest query models.Pagination true "request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/restaurants/deleted [get]
func (h *HandlerV1) URBListDeleted(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "listDeletedBookRestaurants")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.Pagination
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	response, err := h.Service.BookingService().URBListDeleted(
		ctx, &pbb.ListReq{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// List Deleted Attractions
// @Summary List Deleted Attractions
// @Security BearerAuth
// @Description Api for List Deleted Attractions
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/attractions/deleted [get]
func (h *HandlerV1) UABListDeleted(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "ListDeletedBookAttractions")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.Pagination
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	response, err := h.Service.BookingService().UABListDeleted(
		ctx, &pbb.ListReq{
			Limit: body.Limit,
			Offset: (body.Page-1)*body.Limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Update Booked Hotel
// @Summary Update Booked Hotel
// @Security BearerAuth
// @Description Api for Update Booked Hotel
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param models.UpdateBookingReq body models.UpdateBookingReq true "createModel"
// @Success 200 {object} models.BookingRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/hotels [put]
func (h *HandlerV1) UHBUpdate(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "UpdateBookedHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()
	
	var (
		body        models.UpdateBookingReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := h.Service.BookingService().UHBUpdate(ctx, &pbb.GeneralBook{
		Id:                   body.Id.String(),
		UserId:               body.UserId,
		HraId:                body.HraId,
		WillArrive:           body.WillArrive,
		WillLeave:            body.WillLeave,
		NumberOfPeople:       body.NumberOfPeople,
		IsCanceled:           body.IsCanceled,
		Reason:               body.Reason,
		CreatedAt:            time.Now().Format("2006-01-02T15:04:05"),
		UpdatedAt:            time.Now().Format("2006-01-02T15:04:05"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to update booked hotel", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, &models.BookingRes{
		Id:                   uuid.MustParse(response.Id),
		UserId:               response.UserId,
		HraId:                response.HraId,
		WillArrive:           response.WillArrive,
		WillLeave:            response.WillLeave,
		NumberOfPeople:       response.NumberOfPeople,
		IsCanceled:           response.IsCanceled,
		Reason:               response.Reason,
		CreatedAt:            response.CreatedAt,
	})
}

// Update Booked Restaurant
// @Summary Update Booked Restaurant
// @Security BearerAuth
// @Description Api for Update Booked Restaurant
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param models.UpdateBookingReq body models.UpdateBookingReq true "createModel"
// @Success 200 {object} models.BookingRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/restaurants [put]
func (h *HandlerV1) URBUpdate(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "UpdateBookedRestaurant")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.UpdateBookingReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := h.Service.BookingService().URBUpdate(ctx, &pbb.GeneralBook{
		Id:                   body.Id.String(),
		UserId:               body.UserId,
		HraId:                body.HraId,
		WillArrive:           body.WillArrive,
		WillLeave:            body.WillLeave,
		NumberOfPeople:       body.NumberOfPeople,
		IsCanceled:           body.IsCanceled,
		Reason:               body.Reason,
		CreatedAt:            time.Now().Format("2006-01-02T15:04:05"),
		UpdatedAt:            time.Now().Format("2006-01-02T15:04:05"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to update booked hotel", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, &models.BookingRes{
		Id:                   uuid.MustParse(response.Id),
		UserId:               response.UserId,
		HraId:                response.HraId,
		WillArrive:           response.WillArrive,
		WillLeave:            response.WillLeave,
		NumberOfPeople:       response.NumberOfPeople,
		IsCanceled:           response.IsCanceled,
		Reason:               response.Reason,
		CreatedAt:            response.CreatedAt,
	})
}

// Update Booked Attraction
// @Summary Update Booked Attraction
// @Security BearerAuth
// @Description Api for Update Booked Attraction
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param models.UpdateBookingReq body models.UpdateBookingReq true "createModel"
// @Success 200 {object} models.BookingRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/attractions [put]
func (h *HandlerV1) UABUpdate(c *gin.Context) {

	ctx, span := otlp.Start(c, "api", "UpdateBookedAttraction")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.UpdateBookingReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := h.Service.BookingService().UABUpdate(ctx, &pbb.GeneralBook{
		Id:                   body.Id.String(),
		UserId:               body.UserId,
		HraId:                body.HraId,
		WillArrive:           body.WillArrive,
		WillLeave:            body.WillLeave,
		NumberOfPeople:       body.NumberOfPeople,
		IsCanceled:           body.IsCanceled,
		Reason:               body.Reason,
		CreatedAt:            time.Now().Format("2006-01-02T15:04:05"),
		UpdatedAt:            time.Now().Format("2006-01-02T15:04:05"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to update booked hotel", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, &models.BookingRes{
		Id:                   uuid.MustParse(response.Id),
		UserId:               response.UserId,
		HraId:                response.HraId,
		WillArrive:           response.WillArrive,
		WillLeave:            response.WillLeave,
		NumberOfPeople:       response.NumberOfPeople,
		IsCanceled:           response.IsCanceled,
		Reason:               response.Reason,
		CreatedAt:            response.CreatedAt,
	})
}

// Delete Hotel
// @Summary Delete Hotel
// @Security BearerAuth
// @Description Api for Delete Hotel
// @Tags BOOKING_HOTEL
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.StandartError
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/hotels/{id} [delete]
func (h *HandlerV1) UHBDelete(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "DeleteBookedHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")

	response, err := h.Service.BookingService().UHBDelete(
		ctx, &pbb.Id{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to delete booked hotel", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Restaurant
// @Summary Delete Restaurant
// @Security BearerAuth
// @Description Api for Delete Restaurant
// @Tags BOOKING_RESTAURANT
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.StandartError
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/restaurants/{id} [delete]
func (h *HandlerV1) URBDelete(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "DeleteBookedRestaurant")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")

	response, err := h.Service.BookingService().URBDelete(
		ctx, &pbb.Id{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to delete booked restaurant", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Attraction
// @Summary Delete Attraction
// @Security BearerAuth
// @Description Api for Delete Attraction
// @Tags BOOKING_ATTRACTION
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.StandartError
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/booking/attractions/{id} [delete]
func (h *HandlerV1) UABDelete(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "DeleteBookedAttraction")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")

	response, err := h.Service.BookingService().UABDelete(
		ctx, &pbb.Id{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to delete booked attraction", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}