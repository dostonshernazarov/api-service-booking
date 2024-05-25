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

// CREATE ATTRACTION
// @Summary CREATE ATTRACTION
// @Security BearerAuth
// @Description Api for creating attraction
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param Attraction body models.CreateAttraction true "Attraction"
// @Success 200 {object} models.AttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction [POST]
func (h HandlerV1) CreateAttraction(c *gin.Context) {
	var (
		body        models.CreateAttraction
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	
	ctx, span := otlp.Start(c, "api", "CreateAttraction")
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

	owner_id, statusCode := GetIdFromToken(c.Request, h.Config)
	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{
            "error": "Can't get",
        })
        return
    }

	attraction_id := uuid.New().String()
	location_id := uuid.New().String()

	// get images
	var images []*pbe.Image

	for _, bodyImage := range body.Images {
		image_id := uuid.New().String()
		image := pbe.Image{
			ImageId:         image_id,
			EstablishmentId: attraction_id,
			ImageUrl:        bodyImage.ImageUrl,
			Category:        "attraction",
		}

		images = append(images, &image)
	}

	// Format("2006-01-02T15:04:05Z")
	response, err := h.Service.EstablishmentService().CreateAttraction(ctx, &pbe.Attraction{
		AttractionId:   attraction_id,
		OwnerId:        owner_id,
		AttractionName: body.AttractionName,
		Description:    body.Description,
		Rating:         float32(body.Rating),
		ContactNumber:  body.ContactNumber,
		LicenceUrl:     body.LicenceUrl,
		WebsiteUrl:     body.WebsiteUrl,
		Images:         images,
		Location: &pbe.Location{
			LocationId:      location_id,
			EstablishmentId: attraction_id,
			Address:         body.Location.Address,
			Latitude:        float32(body.Location.Latitude),
			Longitude:       float32(body.Location.Longitude),
			Country:         body.Location.Country,
			City:            body.Location.City,
			StateProvince:   body.Location.StateProvince,
			Category:        "attraction",
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

	respModel := models.AttractionModel{
		AttractionId:   response.AttractionId,
		OwnerId:        response.OwnerId,
		AttractionName: response.AttractionName,
		Description:    response.Description,
		Rating:         response.Rating,
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

	c.JSON(http.StatusCreated, respModel)
}

// GET ATTRACTION BY ATTRACTION_ID
// @Summary GET ATTRACTION BY ATTRACTION_ID
// @Security BearerAuth
// @Description Api for getting attraction by attraction_id
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param attraction_id query string true "attraction_id"
// @Success 200 {object} models.AttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction [GET]
func (h HandlerV1) GetAttraction(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "GetAttraction")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	attraction_id := c.Query("attraction_id")

	response, err := h.Service.EstablishmentService().GetAttraction(ctx, &pbe.GetAttractionRequest{
		AttractionId: attraction_id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var respImages []*models.ImageModel

	for _, respImage := range response.Attraction.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.AttractionModel{
		AttractionId:   response.Attraction.AttractionId,
		OwnerId:        response.Attraction.OwnerId,
		AttractionName: response.Attraction.AttractionName,
		Description:    response.Attraction.Description,
		Rating:         response.Attraction.Rating,
		ContactNumber:  response.Attraction.ContactNumber,
		LicenceUrl:     response.Attraction.LicenceUrl,
		WebsiteUrl:     response.Attraction.WebsiteUrl,
		Images:         respImages,
		Location: models.LocationModel{
			LocationId:      response.Attraction.Location.LocationId,
			EstablishmentId: response.Attraction.Location.EstablishmentId,
			Address:         response.Attraction.Location.Address,
			Latitude:        float64(response.Attraction.Location.Latitude),
			Longitude:       float64(response.Attraction.Location.Longitude),
			Country:         response.Attraction.Location.Country,
			City:            response.Attraction.Location.City,
			StateProvince:   response.Attraction.Location.StateProvince,
			CreatedAt:       response.Attraction.Location.CreatedAt,
			UpdatedAt:       response.Attraction.Location.UpdatedAt,
		},
		CreatedAt: response.Attraction.CreatedAt,
		UpdatedAt: response.Attraction.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// LIST ATTRACTIONS BY PAGE AND LIMIT
// @Summary LIST ATTRACTIONS BY PAGE AND LIMIT
// @Security BearerAuth
// @Description Api for listing attractions by page and limit
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.ListAttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction/list [GET]
func (h HandlerV1) ListAttractions(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "ListAttractions")
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

	response, err := h.Service.EstablishmentService().ListAttractions(ctx, &pbe.ListAttractionsRequest{
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

	var respAttractions []*models.AttractionModel

	for _, respAttraction := range response.Attractions {

		var respImages []*models.ImageModel

		for _, respImage := range respAttraction.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		attraction := models.AttractionModel{
			AttractionId:   respAttraction.AttractionId,
			OwnerId:        respAttraction.OwnerId,
			AttractionName: respAttraction.AttractionName,
			Description:    respAttraction.Description,
			Rating:         respAttraction.Rating,
			ContactNumber:  respAttraction.ContactNumber,
			LicenceUrl:     respAttraction.LicenceUrl,
			WebsiteUrl:     respAttraction.WebsiteUrl,
			Images:         respImages,
			Location: models.LocationModel{
				LocationId:      respAttraction.Location.LocationId,
				EstablishmentId: respAttraction.Location.EstablishmentId,
				Address:         respAttraction.Location.Address,
				Latitude:        float64(respAttraction.Location.Latitude),
				Longitude:       float64(respAttraction.Location.Longitude),
				Country:         respAttraction.Location.Country,
				City:            respAttraction.Location.City,
				StateProvince:   respAttraction.Location.StateProvince,
				CreatedAt:       respAttraction.Location.CreatedAt,
				UpdatedAt:       respAttraction.Location.UpdatedAt,
			},
			CreatedAt: respAttraction.CreatedAt,
			UpdatedAt: respAttraction.UpdatedAt,
		}

		respAttractions = append(respAttractions, &attraction)
	}

	listModel := models.ListAttractionModel{
		Attractions: respAttractions,
		Count:       response.Overall,
	}

	c.JSON(200, listModel)
}

// UPDATE ATTRACTION
// @Summary UPDATE ATTRACTION
// @Security BearerAuth
// @Description Api for updating attraction by attraction_id
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param attraction_id query string true "attraction_id"
// @Param UpdatingAttraction body models.UpdateAttraction true "UpdatingAttraction"
// @Success 200 {object} models.AttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction [PUT]
func (h HandlerV1) UpdateAttraction(c *gin.Context) {
	var (
		body        models.Attraction
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "UpdateAttraction")
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

	attraction_id := c.Query("attraction_id")

	response, err := h.Service.EstablishmentService().UpdateAttraction(ctx, &pbe.UpdateAttractionRequest{
		Attraction: &pbe.Attraction{
			AttractionId:   attraction_id,
			AttractionName: body.AttractionName,
			Description:    body.Description,
			Rating:         float32(body.Rating),
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

	for _, respImage := range response.Attraction.Images {
		image := models.ImageModel{
			ImageId:         respImage.ImageId,
			EstablishmentId: respImage.EstablishmentId,
			ImageUrl:        respImage.ImageUrl,
			CreatedAt:       respImage.CreatedAt,
			UpdatedAt:       respImage.UpdatedAt,
		}
		respImages = append(respImages, &image)
	}

	respModel := models.AttractionModel{
		AttractionId:   response.Attraction.AttractionId,
		OwnerId:        response.Attraction.OwnerId,
		AttractionName: response.Attraction.AttractionName,
		Description:    response.Attraction.Description,
		Rating:         response.Attraction.Rating,
		ContactNumber:  response.Attraction.ContactNumber,
		LicenceUrl:     response.Attraction.LicenceUrl,
		WebsiteUrl:     response.Attraction.WebsiteUrl,
		Images:         respImages,
		Location: models.LocationModel{
			LocationId:      response.Attraction.Location.LocationId,
			EstablishmentId: response.Attraction.Location.EstablishmentId,
			Address:         response.Attraction.Location.Address,
			Latitude:        float64(response.Attraction.Location.Latitude),
			Longitude:       float64(response.Attraction.Location.Longitude),
			Country:         response.Attraction.Location.Country,
			City:            response.Attraction.Location.City,
			StateProvince:   response.Attraction.Location.StateProvince,
			CreatedAt:       response.Attraction.Location.CreatedAt,
			UpdatedAt:       response.Attraction.Location.UpdatedAt,
		},
		CreatedAt: response.Attraction.CreatedAt,
		UpdatedAt: response.Attraction.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// DELETE ATTRACTION BY ATTRACTION_ID
// @Summary DELETE ATTRACTION BY ATTRACTION_ID
// @Security BearerAuth
// @Description Api for deleting attraction by attraction_id
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param attraction_id query string true "attraction_id"
// @Success 200 {object} models.DeleteResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction [DELETE]
func (h HandlerV1) DeleteAttraction(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	ctx, span := otlp.Start(c, "api", "DeleteAttraction")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	attraction_id := c.Query("attraction_id")

	response, err := h.Service.EstablishmentService().DeleteAttraction(ctx, &pbe.DeleteAttractionRequest{
		AttractionId: attraction_id,
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

// LIST ATTRACTIONS BY PAGE, LIMIT, COUNTRY, CITY AND STATE_PROVINCE
// @Summary LIST ATTRACTIONS BY PAGE, LIMIT, COUNTRY, CITY AND STATE_PROVINCE
// @Security BearerAuth
// @Description Api for listing attractions by page, limit, country, city and state_province
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Param request query models.FieldValuesByLocation true "request"
// @Success 200 {object} models.ListAttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction/listlocation [GET]  
func (h HandlerV1) ListAttractionsByLocation(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	ctx, span := otlp.Start(c, "api", "ListAttractions")
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
	state_province := c.Query("province")

	response, err := h.Service.EstablishmentService().ListAttractionsByLocation(ctx, &pbe.ListAttractionsByLocationRequest{
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

	var respAttractions []*models.AttractionModel

	for _, respAttraction := range response.Attractions {

		var respImages []*models.ImageModel

		for _, respImage := range respAttraction.Images {
			image := models.ImageModel{
				ImageId:         respImage.ImageId,
				EstablishmentId: respImage.EstablishmentId,
				ImageUrl:        respImage.ImageUrl,
				CreatedAt:       respImage.CreatedAt,
				UpdatedAt:       respImage.UpdatedAt,
			}
			respImages = append(respImages, &image)
		}

		attraction := models.AttractionModel{
			AttractionId:   respAttraction.AttractionId,
			OwnerId:        respAttraction.OwnerId,
			AttractionName: respAttraction.AttractionName,
			Description:    respAttraction.Description,
			Rating:         respAttraction.Rating,
			ContactNumber:  respAttraction.ContactNumber,
			LicenceUrl:     respAttraction.LicenceUrl,
			WebsiteUrl:     respAttraction.WebsiteUrl,
			Images:         respImages,
			Location: models.LocationModel{
				LocationId:      respAttraction.Location.LocationId,
				EstablishmentId: respAttraction.Location.EstablishmentId,
				Address:         respAttraction.Location.Address,
				Latitude:        float64(respAttraction.Location.Latitude),
				Longitude:       float64(respAttraction.Location.Longitude),
				Country:         respAttraction.Location.Country,
				City:            respAttraction.Location.City,
				StateProvince:   respAttraction.Location.StateProvince,
				CreatedAt:       respAttraction.Location.CreatedAt,
				UpdatedAt:       respAttraction.Location.UpdatedAt,
			},
			CreatedAt: respAttraction.CreatedAt,
			UpdatedAt: respAttraction.UpdatedAt,
		}

		respAttractions = append(respAttractions, &attraction)
	}

	respModel := models.ListAttractionModel{
		Attractions: respAttractions,
		Count:     uint64(response.Count),
	}

	c.JSON(200, respModel)
}