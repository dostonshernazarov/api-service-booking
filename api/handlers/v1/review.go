package v1

import (
	"Booking/api-service-booking/api/models"
	pb "Booking/api-service-booking/genproto/establishment-proto"
	"Booking/api-service-booking/internal/pkg/otlp"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// CREATE REVIEW
// @Summary CREATE REVIEW
// @Security BearerAuth
// @Description Api for creating review
// @Tags REVIEW
// @Accept json
// @Produce json
// @Param establishment_id query string true "establishment_id"
// @Param user_id query string true "user_id"
// @Param Review body models.CreateReview true "Review"
// @Success 200 {object} models.ReviewModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/review/create [POST]
func (h HandlerV1) CreateReview(c *gin.Context) {
	var (
		body        models.CreateReview
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "CreateReview")
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

	establishment_id := c.Query("establishment_id")
	user_id := c.Query("user_id")

	review_id := uuid.New().String()

	response, err := h.Service.EstablishmentService().CreateReview(ctx, &pb.CreateReviewRequest{
		Review: &pb.Review{
			ReviewId:        review_id,
			EstablishmentId: establishment_id,
			UserId:          user_id,
			Rating:          float32(body.Rating),
			Comment:         body.Comment,
		},
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	respModel := models.ReviewModel{
		ReviewId:        response.Review.ReviewId,
		EstablishmentId: response.Review.EstablishmentId,
		UserId:          response.Review.UserId,
		Rating:          float64(response.Review.Rating),
		Comment:         response.Review.Comment,
		CreatedAt:       response.Review.CreatedAt,
		UpdatedAt:       response.Review.UpdatedAt,
	}

	c.JSON(200, respModel)
}

// LIST REVIEWS BY ESTABLISHMENT_ID
// @Summary LIST REVIEWS BY ESTABLISHMENT_ID
// @Security BearerAuth
// @Description Api for listing reviews by establishment_id
// @Tags REVIEW
// @Accept json
// @Produce json
// @Param establishment_id query string true "establishment_id"
// @Success 200 {object} models.ListReviews
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/review/list [GET]
func (h HandlerV1) ListReviews(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "ListReviews")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	establishment_id := c.Query("establishment_id")

	response, err := h.Service.EstablishmentService().ListReviews(ctx, &pb.ListReviewsRequest{
		EstablishmentId: establishment_id,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var reviews []*models.ReviewModel

	for _, respReview := range response.Reviews {
		review := models.ReviewModel{
			ReviewId:        respReview.ReviewId,
			EstablishmentId: respReview.EstablishmentId,
			UserId:          respReview.UserId,
			Rating:          float64(respReview.Rating),
			Comment:         respReview.Comment,
			CreatedAt:       respReview.CreatedAt,
			UpdatedAt:       respReview.UpdatedAt,
		}

		reviews = append(reviews, &review)
	}

	respModel := models.ListReviews{
		Reviews: reviews,
		Count:   response.Count,
	}

	c.JSON(200, respModel)
}

// DELETE REVIEW BY REVIEW_ID
// @Summary DELETE REVIEW BY REVIEW_ID
// @Security BearerAuth
// @Description Api for deleting review by review_id
// @Tags REVIEW
// @Accept json
// @Produce json
// @Param review_id query string true "review_id"
// @Success 200 {object} models.DeleteResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/review/delete [DELETE]
func (h HandlerV1) DeleteReview(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, span := otlp.Start(c, "api", "DeleteReview")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)
	defer span.End()

	review_id := c.Query("review_id")

	response, err := h.Service.EstablishmentService().DeleteReview(ctx, &pb.DeleteReviewRequest{
		ReviewId: review_id,
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
	}

	c.JSON(200, gin.H{
		"message": "successfully deleted",
	})
}

