package v1

import (
	"Booking/api-service-booking/api/models"
	pbe "Booking/api-service-booking/genproto/establishment-proto"
	"Booking/api-service-booking/internal/pkg/otlp"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// CREATE HOTEL
// @Summary CREATE HOTEL
// @Security BearerAuth
// @Description Api for creating hotel
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param Hotel body models.CreateHotel true "Hotel"
// @Success 200 {object} models.HotelModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel [POST]
func (h HandlerV1) CreateHotel(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "CreateHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	var (
		body        models.CreateHotel
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	owner_id, statusCode := GetIdFromToken(c.Request, h.Config)
	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{
            "error": "Can't get",
        })
        return
    }


	hotel_id := uuid.New().String()
	location_id := uuid.New().String()

	// get images
	var images []*pbe.Image

	for _, bodyImage := range body.Images {
		image_id := uuid.New().String()
		image := pbe.Image{
			ImageId:         image_id,
			EstablishmentId: hotel_id,
			ImageUrl:        bodyImage.ImageUrl,
			Category:        "hotel",
		}

		images = append(images, &image)
	}

	// Format("2006-01-02T15:04:05Z")
	response, err := h.Service.EstablishmentService().CreateHotel(ctx, &pbe.Hotel{
		HotelId:       hotel_id,
		OwnerId:       owner_id,
		HotelName:     body.HotelName,
		Description:   body.Description,
		Rating:        float32(body.Rating),
		ContactNumber: body.ContactNumber,
		LicenceUrl:    body.LicenceUrl,
		WebsiteUrl:    body.WebsiteUrl,
		Images:        images,
		Location: &pbe.Location{
			LocationId:      location_id,
			EstablishmentId: hotel_id,
			Address:         body.Location.Address,
			Latitude:        float32(body.Location.Latitude),
			Longitude:       float32(body.Location.Longitude),
			Country:         body.Location.Country,
			City:            body.Location.City,
			StateProvince:   body.Location.StateProvince,
			Category:        "hotel",
		},
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var respImages []*models.ImageModel

	for _, respImage := range response.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.HotelModel{
		HotelId:       response.HotelId,
		OwnerId:       response.OwnerId,
		HotelName:     response.HotelName,
		Description:   response.Description,
		Rating:        response.Rating,
		ContactNumber: response.ContactNumber,
		LicenceUrl:    response.LicenceUrl,
		WebsiteUrl:    response.WebsiteUrl,
		Images:        respImages,
		Location: models.LocationModel{
			LocationId:      response.Location.LocationId,
			EstablishmentId: response.Location.EstablishmentId,
			Address:         response.Location.Address,
			Latitude:        float64(response.Location.Latitude),
			Longitude:       float64(response.Location.Longitude),
			Country:         response.Location.Country,
			City:            response.Location.City,
			StateProvince:   response.Location.StateProvince,
			CreatedAt:       response.Location.CreatedAt,
			UpdatedAt:       response.Location.UpdatedAt,
		},
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}

	c.JSON(http.StatusCreated, respModel)
}

// GET HOTEL BY HOTEL_ID
// @Summary GET HOTEL BY HOTEL_ID
// @Security BearerAuth
// @Description Api for getting hotel by hotel_id
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param hotel_id query string true "hotel_id"
// @Success 200 {object} models.HotelModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel [GET]
func (h HandlerV1) GetHotel(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "GetHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	hotel_id := c.Query("hotel_id")

	response, err := h.Service.EstablishmentService().GetHotel(ctx, &pbe.GetHotelRequest{
		HotelId: hotel_id,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respImages []*models.ImageModel

	for _, respImage := range response.Hotel.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.HotelModel{
		HotelId:       response.Hotel.HotelId,
		OwnerId:       response.Hotel.OwnerId,
		HotelName:     response.Hotel.HotelName,
		Description:   response.Hotel.Description,
		Rating:        response.Hotel.Rating,
		ContactNumber: response.Hotel.ContactNumber,
		LicenceUrl:    response.Hotel.LicenceUrl,
		WebsiteUrl:    response.Hotel.WebsiteUrl,
		Images:        respImages,
		Location: models.LocationModel{
			LocationId:      response.Hotel.Location.LocationId,
			EstablishmentId: response.Hotel.Location.EstablishmentId,
			Address:         response.Hotel.Location.Address,
			Latitude:        float64(response.Hotel.Location.Latitude),
			Longitude:       float64(response.Hotel.Location.Longitude),
			Country:         response.Hotel.Location.Country,
			City:            response.Hotel.Location.City,
			StateProvince:   response.Hotel.Location.StateProvince,
			CreatedAt:       response.Hotel.Location.CreatedAt,
			UpdatedAt:       response.Hotel.Location.UpdatedAt,
		},
		CreatedAt: response.Hotel.CreatedAt,
		UpdatedAt: response.Hotel.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// LIST HOTELS BY PAGE AND LIMIT
// @Summary LIST HOTELS BY PAGE AND LIMIT
// @Security BearerAuth
// @Description Api for listing hotels by page and limit
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.ListHotelsModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel/list [GET]
func (h HandlerV1) ListHotels(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "ListHotels")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	limit := c.Query("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	offset := (pageInt - 1) * limitInt

	response, err := h.Service.EstablishmentService().ListHotels(ctx, &pbe.ListHotelsRequest{
		Offset: int64(offset),
		Limit:  int64(limitInt),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respHotels []*models.HotelModel

	for _, respHotel := range response.Hotels {

		var respImages []*models.ImageModel

		for _, respImage := range respHotel.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		hotel := models.HotelModel{
			HotelId:       respHotel.HotelId,
			OwnerId:       respHotel.OwnerId,
			HotelName:     respHotel.HotelName,
			Description:   respHotel.Description,
			Rating:        respHotel.Rating,
			ContactNumber: respHotel.ContactNumber,
			LicenceUrl:    respHotel.LicenceUrl,
			WebsiteUrl:    respHotel.WebsiteUrl,
			Images:        respImages,
			Location: models.LocationModel{
				LocationId:      respHotel.Location.LocationId,
				EstablishmentId: respHotel.Location.EstablishmentId,
				Address:         respHotel.Location.Address,
				Latitude:        float64(respHotel.Location.Latitude),
				Longitude:       float64(respHotel.Location.Longitude),
				Country:         respHotel.Location.Country,
				City:            respHotel.Location.City,
				StateProvince:   respHotel.Location.StateProvince,
				CreatedAt:       respHotel.Location.CreatedAt,
				UpdatedAt:       respHotel.Location.UpdatedAt,
			},
			CreatedAt: respHotel.CreatedAt,
			UpdatedAt: respHotel.UpdatedAt,
		}

		respHotels = append(respHotels, &hotel)
	}

	listModel := models.ListHotelsModel{
		Hotels:  respHotels,
		Count: response.Overall,
	}

	c.JSON(200, listModel)
}

// UPDATE HOTEL
// @Summary UPDATE HOTEL
// @Security BearerAuth
// @Description Api for updating hotel by hotel_id
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param hotel_id query string true "hotel_id"
// @Param UpdatingHotel body models.UpdateHotel true "UpdatingHotel"
// @Success 200 {object} models.HotelModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel [PUT]
func (h HandlerV1) UpdateHotel(c *gin.Context) {

	var (
		body        models.CreateHotel
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "UpdateHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	hotel_id := c.Query("hotel_id")

	response, err := h.Service.EstablishmentService().UpdateHotel(ctx, &pbe.UpdateHotelRequest{
		Hotel: &pbe.Hotel{
			HotelId:       hotel_id,
			HotelName:     body.HotelName,
			Description:   body.Description,
			Rating:        float32(body.Rating),
			ContactNumber: body.ContactNumber,
			LicenceUrl:    body.LicenceUrl,
			WebsiteUrl:    body.WebsiteUrl,
			Location: &pbe.Location{
				Address:       body.Location.Address,
				Latitude:      float32(body.Location.Latitude),
				Longitude:     float32(body.Location.Longitude),
				Country:       body.Location.Country,
				City:          body.Location.City,
				StateProvince: body.Location.StateProvince,
			},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respImages []*models.ImageModel

	for _, respImage := range response.Hotel.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.HotelModel{
		HotelId:       response.Hotel.HotelId,
		OwnerId:       response.Hotel.OwnerId,
		HotelName:     response.Hotel.HotelName,
		Description:   response.Hotel.Description,
		Rating:        response.Hotel.Rating,
		ContactNumber: response.Hotel.ContactNumber,
		LicenceUrl:    response.Hotel.LicenceUrl,
		WebsiteUrl:    response.Hotel.WebsiteUrl,
		Images:        respImages,
		Location: models.LocationModel{
			LocationId:      response.Hotel.Location.LocationId,
			EstablishmentId: response.Hotel.Location.EstablishmentId,
			Address:         response.Hotel.Location.Address,
			Latitude:        float64(response.Hotel.Location.Latitude),
			Longitude:       float64(response.Hotel.Location.Longitude),
			Country:         response.Hotel.Location.Country,
			City:            response.Hotel.Location.City,
			StateProvince:   response.Hotel.Location.StateProvince,
			CreatedAt:       response.Hotel.Location.CreatedAt,
			UpdatedAt:       response.Hotel.Location.UpdatedAt,
		},
		CreatedAt: response.Hotel.CreatedAt,
		UpdatedAt: response.Hotel.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// DELETE HOTEL BY HOTEL_ID
// @Summary DELETE HOTEL BY HOTEL_ID
// @Security BearerAuth
// @Description Api for deleting hotel by hotel_id
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param hotel_id query string true "hotel_id"
// @Success 200 {object} models.DeleteResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel [DELETE]
func (h HandlerV1) DeleteHotel(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "DeleteHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	hotel_id := c.Query("hotel_id")

	response, err := h.Service.EstablishmentService().DeleteHotel(ctx, &pbe.DeleteHotelRequest{
		HotelId: hotel_id,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	if !response.Success {
		c.JSON(404, gin.H{
			"error": "not deleted",
		})
		h.Logger.Error("not deleted")
		return
	}

	c.JSON(200, gin.H{
		"message": "successfuly deleted",
	})
}

// LIST HOTELS BY PAGE, LIMIT, COUNTRY, CITY AND STATE_PROVINCE
// @Summary LIST HOTELS BY PAGE, LIMIT, COUNTRY, CITY AND STATE_PROVINCE
// @Security BearerAuth
// @Description Api for listing hotels by page, limit, country, city and state_province
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Param request query models.FieldValuesByLocation true "request"
// @Success 200 {object} models.ListHotelsModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel/listlocation [GET]
func (h HandlerV1) ListHotelsByLocation(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "ListHotels")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	limit := c.Query("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	offset := (pageInt - 1) * limitInt

	country := c.Query("country")
	city := c.Query("city")
	state_province := c.Query("state_province")

	response, err := h.Service.EstablishmentService().ListHotelsByLocation(ctx, &pbe.ListHotelsByLocationRequest{
		Offset:        uint64(offset),
		Limit:         uint64(limitInt),
		Country:       country,
		City:          city,
		StateProvince: state_province,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respHotels []*models.HotelModel

	for _, respHotel := range response.Hotels {

		var respImages []*models.ImageModel

		for _, respImage := range respHotel.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		hotel := models.HotelModel{
			HotelId:       respHotel.HotelId,
			OwnerId:       respHotel.OwnerId,
			HotelName:     respHotel.HotelName,
			Description:   respHotel.Description,
			Rating:        respHotel.Rating,
			ContactNumber: respHotel.ContactNumber,
			LicenceUrl:    respHotel.LicenceUrl,
			WebsiteUrl:    respHotel.WebsiteUrl,
			Images:        respImages,
			Location: models.LocationModel{
				LocationId:      respHotel.Location.LocationId,
				EstablishmentId: respHotel.Location.EstablishmentId,
				Address:         respHotel.Location.Address,
				Latitude:        float64(respHotel.Location.Latitude),
				Longitude:       float64(respHotel.Location.Longitude),
				Country:         respHotel.Location.Country,
				City:            respHotel.Location.City,
				StateProvince:   respHotel.Location.StateProvince,
				CreatedAt:       respHotel.Location.CreatedAt,
				UpdatedAt:       respHotel.Location.UpdatedAt,
			},
			CreatedAt: respHotel.CreatedAt,
			UpdatedAt: respHotel.UpdatedAt,
		}

		respHotels = append(respHotels, &hotel)
	}

	respModel := models.ListHotelsModel{
		Hotels:  respHotels,
		Count: uint64(response.Count),
	}

	c.JSON(200, respModel)
}

// FIND HOTELS BY NAME
// @Summary FIND HOTELS BY NAME
// @Security BearerAuth
// @Description Api for listing hotels by name
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Success 200 {object} models.ListHotelsModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/hotel/find [GET]
func (h HandlerV1) FindHotelsByName(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "ListHotels")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	name := c.Query("name")
	

	response, err := h.Service.EstablishmentService().FindHotelsByName(ctx, &pbe.FindHotelsByNameRequest{
		Name: name,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respHotels []*models.HotelModel

	for _, respHotel := range response.Hotels {

		var respImages []*models.ImageModel

		for _, respImage := range respHotel.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		hotel := models.HotelModel{
			HotelId:       respHotel.HotelId,
			OwnerId:       respHotel.OwnerId,
			HotelName:     respHotel.HotelName,
			Description:   respHotel.Description,
			Rating:        respHotel.Rating,
			ContactNumber: respHotel.ContactNumber,
			LicenceUrl:    respHotel.LicenceUrl,
			WebsiteUrl:    respHotel.WebsiteUrl,
			Images:        respImages,
			Location: models.LocationModel{
				LocationId:      respHotel.Location.LocationId,
				EstablishmentId: respHotel.Location.EstablishmentId,
				Address:         respHotel.Location.Address,
				Latitude:        float64(respHotel.Location.Latitude),
				Longitude:       float64(respHotel.Location.Longitude),
				Country:         respHotel.Location.Country,
				City:            respHotel.Location.City,
				StateProvince:   respHotel.Location.StateProvince,
				CreatedAt:       respHotel.Location.CreatedAt,
				UpdatedAt:       respHotel.Location.UpdatedAt,
			},
			CreatedAt: respHotel.CreatedAt,
			UpdatedAt: respHotel.UpdatedAt,
		}

		respHotels = append(respHotels, &hotel)
	}

	listModel := models.ListHotelsModel{
		Hotels:  respHotels,
		Count: response.Count,
	}

	c.JSON(200, listModel)
}