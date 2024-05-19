package v1

import (
	"Booking/api-service-booking/api/models"
	pbe "Booking/api-service-booking/genproto/establishment-proto"
	"Booking/api-service-booking/internal/pkg/otlp"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// CREATE RESTAURANT
// @Summary CREATE RESTAURANT
// @Security BearerAuth
// @Description Api for creating restaurant
// @Tags RESTAURANT
// @Accept json
// @Produce json
// @Param owner_id query string true "owner_id"
// @Param Restaurant body models.CreateRestaurant true "Restaurant"
// @Success 200 {object} models.RestaurantModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/restaurant [POST]
func (h HandlerV1) CreateRestaurant(c *gin.Context) {
	var (
		body        models.CreateRestaurant
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "CreateRestaurant")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	owner_id := c.Query("owner_id")

	restaurant_id := uuid.New().String()
	location_id := uuid.New().String()

	// get images
	var images []*pbe.Image

	for _, bodyImage := range body.Images {
		image_id := uuid.New().String()
		image := pbe.Image{
			ImageId:         image_id,
			EstablishmentId: restaurant_id,
			ImageUrl:        bodyImage.ImageUrl,
			Category:        "restaurant",
		}

		images = append(images, &image)
	}

	// Format("2006-01-02T15:04:05Z")
	response, err := h.Service.EstablishmentService().CreateRestaurant(ctx, &pbe.Restaurant{
		RestaurantId:   restaurant_id,
		OwnerId:        owner_id,
		RestaurantName: body.RestaurantName,
		Description:    body.Description,
		Rating:         float32(body.Rating),
		OpeningHours:   body.OpeningHours,
		ContactNumber:  body.ContactNumber,
		LicenceUrl:     body.LicenceUrl,
		WebsiteUrl:     body.WebsiteUrl,
		Images:         images,
		Location: &pbe.Location{
			LocationId:      location_id,
			EstablishmentId: restaurant_id,
			Address:         body.Location.Address,
			Latitude:        float32(body.Location.Latitude),
			Longitude:       float32(body.Location.Longitude),
			Country:         body.Location.Country,
			City:            body.Location.City,
			StateProvince:   body.Location.StateProvince,
			Category:        "restaurant",
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

	respModel := models.RestaurantModel{
		RestaurantId:   response.RestaurantId,
		OwnerId:        response.OwnerId,
		RestaurantName: response.RestaurantName,
		Description:    response.Description,
		Rating:         response.Rating,
		OpeningHours:   response.OpeningHours,
		ContactNumber:  response.ContactNumber,
		LicenceUrl:     response.LicenceUrl,
		WebsiteUrl:     response.WebsiteUrl,
		Images:         respImages,
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

	c.JSON(200, respModel)
}

// GET RESTAURANT BY RESTAURANT_ID
// @Summary GET RESTAURANT BY RESTAURANT_ID
// @Security BearerAuth
// @Description Api for getting restaurant by restaurant_id
// @Tags RESTAURANT
// @Accept json
// @Produce json
// @Param restaurant_id query string true "restaurant_id"
// @Success 200 {object} models.RestaurantModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/restaurant [GET]
func (h HandlerV1) GetRestaurant(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "GetHotel")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	restaurant_id := c.Query("restaurant_id")

	response, err := h.Service.EstablishmentService().GetRestaurant(ctx, &pbe.GetRestaurantRequest{
		RestaurantId: restaurant_id,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respImages []*models.ImageModel

	for _, respImage := range response.Restaurant.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.RestaurantModel{
		RestaurantId:   response.Restaurant.RestaurantId,
		OwnerId:        response.Restaurant.OwnerId,
		RestaurantName: response.Restaurant.RestaurantName,
		Description:    response.Restaurant.Description,
		Rating:         response.Restaurant.Rating,
		OpeningHours:   response.Restaurant.OpeningHours,
		ContactNumber:  response.Restaurant.ContactNumber,
		LicenceUrl:     response.Restaurant.LicenceUrl,
		WebsiteUrl:     response.Restaurant.WebsiteUrl,
		Images:         respImages,
		Location: models.LocationModel{
			LocationId:      response.Restaurant.Location.LocationId,
			EstablishmentId: response.Restaurant.Location.EstablishmentId,
			Address:         response.Restaurant.Location.Address,
			Latitude:        float64(response.Restaurant.Location.Latitude),
			Longitude:       float64(response.Restaurant.Location.Longitude),
			Country:         response.Restaurant.Location.Country,
			City:            response.Restaurant.Location.City,
			StateProvince:   response.Restaurant.Location.StateProvince,
			CreatedAt:       response.Restaurant.Location.CreatedAt,
			UpdatedAt:       response.Restaurant.Location.UpdatedAt,
		},
		CreatedAt: response.Restaurant.CreatedAt,
		UpdatedAt: response.Restaurant.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// LIST RESTAURANTS BY PAGE AND LIMIT
// @Summary LIST RESTAURANTS BY PAGE AND LIMIT
// @Security BearerAuth
// @Description Api for listing restaurants by page and limit
// @Tags RESTAURANT
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.ListRestaurantsModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/restaurant/list [GET]
func (h HandlerV1) ListRestaurants(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "ListRestaurants")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

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

	response, err := h.Service.EstablishmentService().ListRestaurants(ctx, &pbe.ListRestaurantsRequest{
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

	var respRestaurants []*models.RestaurantModel

	for _, respRestaurant := range response.Restaurants {

		var respImages []*models.ImageModel

		for _, respImage := range respRestaurant.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		restaurant := models.RestaurantModel{
			RestaurantId:   respRestaurant.RestaurantId,
			OwnerId:        respRestaurant.OwnerId,
			RestaurantName: respRestaurant.RestaurantName,
			Description:    respRestaurant.Description,
			Rating:         respRestaurant.Rating,
			OpeningHours:   respRestaurant.OpeningHours,
			ContactNumber:  respRestaurant.ContactNumber,
			LicenceUrl:     respRestaurant.LicenceUrl,
			WebsiteUrl:     respRestaurant.WebsiteUrl,
			Images:         respImages,
			Location: models.LocationModel{
				LocationId:      respRestaurant.Location.LocationId,
				EstablishmentId: respRestaurant.Location.EstablishmentId,
				Address:         respRestaurant.Location.Address,
				Latitude:        float64(respRestaurant.Location.Latitude),
				Longitude:       float64(respRestaurant.Location.Longitude),
				Country:         respRestaurant.Location.Country,
				City:            respRestaurant.Location.City,
				StateProvince:   respRestaurant.Location.StateProvince,
				CreatedAt:       respRestaurant.Location.CreatedAt,
				UpdatedAt:       respRestaurant.Location.UpdatedAt,
			},
			CreatedAt: respRestaurant.CreatedAt,
			UpdatedAt: respRestaurant.UpdatedAt,
		}

		respRestaurants = append(respRestaurants, &restaurant)
	}

	listModel := models.ListRestaurantsModel{
		Restaurants: respRestaurants,
		Count:     response.Overall,
	}

	c.JSON(200, listModel)
}

// UPDATE RESTAURANT
// @Summary UPDATE RESTAURANT
// @Security BearerAuth
// @Description Api for updating restaurant by restaurant_id
// @Tags RESTAURANT
// @Accept json
// @Produce json
// @Param restaurant_id query string true "restaurant_id"
// @Param UpdatingRestaurant body models.UpdateRestaurant true "UpdatingRestaurant"
// @Success 200 {object} models.RestaurantModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/restaurant [PUT]
func (h HandlerV1) UpdateRestaurant(c *gin.Context) {
	var (
		body        models.CreateRestaurant
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "UpdateRestaurant")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	restaurant_id := c.Query("restaurant_id")

	response, err := h.Service.EstablishmentService().UpdateRestaurant(ctx, &pbe.UpdateRestaurantRequest{
		Restaurant: &pbe.Restaurant{
			RestaurantId:   restaurant_id,
			RestaurantName: body.RestaurantName,
			Description:    body.Description,
			Rating:         float32(body.Rating),
			OpeningHours:   body.OpeningHours,
			ContactNumber:  body.ContactNumber,
			LicenceUrl:     body.LicenceUrl,
			WebsiteUrl:     body.WebsiteUrl,
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

	for _, respImage := range response.Restaurant.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.RestaurantModel{
		RestaurantId:   response.Restaurant.RestaurantId,
		OwnerId:        response.Restaurant.OwnerId,
		RestaurantName: response.Restaurant.RestaurantName,
		Description:    response.Restaurant.Description,
		Rating:         response.Restaurant.Rating,
		OpeningHours:   response.Restaurant.OpeningHours,
		ContactNumber:  response.Restaurant.ContactNumber,
		LicenceUrl:     response.Restaurant.LicenceUrl,
		WebsiteUrl:     response.Restaurant.WebsiteUrl,
		Images:         respImages,
		Location: models.LocationModel{
			LocationId:      response.Restaurant.Location.LocationId,
			EstablishmentId: response.Restaurant.Location.EstablishmentId,
			Address:         response.Restaurant.Location.Address,
			Latitude:        float64(response.Restaurant.Location.Latitude),
			Longitude:       float64(response.Restaurant.Location.Longitude),
			Country:         response.Restaurant.Location.Country,
			City:            response.Restaurant.Location.City,
			StateProvince:   response.Restaurant.Location.StateProvince,
			CreatedAt:       response.Restaurant.Location.CreatedAt,
			UpdatedAt:       response.Restaurant.Location.UpdatedAt,
		},
		CreatedAt: response.Restaurant.CreatedAt,
		UpdatedAt: response.Restaurant.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// DELETE RESTAURANT BY RESTAURANT_ID
// @Summary DELETE RESTAURANT BY RESTAURANT_ID
// @Security BearerAuth
// @Description Api for deleting restaurant by restaurant_id
// @Tags RESTAURANT
// @Accept json
// @Produce json
// @Param restaurant_id query string true "restaurant_id"
// @Success 200 {object} models.DeleteResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/restaurant [DELETE]
func (h HandlerV1) DeleteRestaurant(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "DeleteRestaurant")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	restaurant_id := c.Query("restaurant_id")

	response, err := h.Service.EstablishmentService().DeleteRestaurant(ctx, &pbe.DeleteRestaurantRequest{
		RestaurantId: restaurant_id,
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

// LIST RESTAURANTS BY PAGE, LIMIT, COUNTRY, CITY AND STATE_PROVINCE
// @Summary LIST RESTAURANTS BY PAGE, LIMIT, COUNTRY, CITY AND STATE_PROVINCE
// @Security BearerAuth
// @Description Api for listing restaurants by page, limit, country, city and state_province
// @Tags RESTAURANT
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Param country query string true "country"
// @Param city query string true "city"
// @Param state_province query string true "state_province"
// @Success 200 {object} models.ListRestaurantsModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/restaurant/listlocation [GET]
func (h HandlerV1) ListRestaurantsByLocation(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "ListRestaurants")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

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

	response, err := h.Service.EstablishmentService().ListRestaurantsByLocation(ctx, &pbe.ListRestaurantsByLocationRequest{
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

	var respRestaurants []*models.RestaurantModel

	for _, respRestaurant := range response.Restaurants {

		var respImages []*models.ImageModel

		for _, respImage := range respRestaurant.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		restaurant := models.RestaurantModel{
			RestaurantId:   respRestaurant.RestaurantId,
			OwnerId:        respRestaurant.OwnerId,
			RestaurantName: respRestaurant.RestaurantName,
			Description:    respRestaurant.Description,
			Rating:         respRestaurant.Rating,
			OpeningHours:   respRestaurant.OpeningHours,
			ContactNumber:  respRestaurant.ContactNumber,
			LicenceUrl:     respRestaurant.LicenceUrl,
			WebsiteUrl:     respRestaurant.WebsiteUrl,
			Images:         respImages,
			Location: models.LocationModel{
				LocationId:      respRestaurant.Location.LocationId,
				EstablishmentId: respRestaurant.Location.EstablishmentId,
				Address:         respRestaurant.Location.Address,
				Latitude:        float64(respRestaurant.Location.Latitude),
				Longitude:       float64(respRestaurant.Location.Longitude),
				Country:         respRestaurant.Location.Country,
				City:            respRestaurant.Location.City,
				StateProvince:   respRestaurant.Location.StateProvince,
				CreatedAt:       respRestaurant.Location.CreatedAt,
				UpdatedAt:       respRestaurant.Location.UpdatedAt,
			},
			CreatedAt: respRestaurant.CreatedAt,
			UpdatedAt: respRestaurant.UpdatedAt,
		}

		respRestaurants = append(respRestaurants, &restaurant)
	}

	respModel := models.ListRestaurantsModel{
		Restaurants: respRestaurants,
		Count:     uint64(response.Count),
	}

	c.JSON(200, respModel)
}