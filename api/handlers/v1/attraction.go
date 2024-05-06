package v1

import (
	"Booking/api_establishment_booking/api/models"
	pbe "Booking/api_establishment_booking/genproto/establishment-proto"
	"Booking/api_establishment_booking/internal/pkg/otlp"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// CREATE ATTRACTION
// @Summary CREATE ATTRACTION
// @Description Api for creating attraction
// @Tags ATTRACTION
// @Accept json
// @Produce json
// @Param Attraction body models.CreateAttraction true "Attraction"
// @Param owner_id query string true "owner_id"
// @Success 200 {object} models.AttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction/create [POST]
func (h HandlerV1) CreateAttraction(c *gin.Context) {
	var (
		body        models.Attraction
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "CreateAttraction")
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

	c.JSON(200, respModel)
}
