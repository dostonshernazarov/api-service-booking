package v1

import (
	"Booking/api-service-booking/api/models"
	pb "Booking/api-service-booking/genproto/establishment-proto"
	"Booking/api-service-booking/internal/pkg/otlp"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// ADD ESTABLISHMENT TO FAVOURITES
// @Summary ADD ESTABLISHMENT TO FAVOURITES
// @Description Api for adding establishment to favourites
// @Tags FAVOURITE
// @Accept json
// @Produce json
// @Param establishment_id query string true "establishment_id"
// @Param user_id query string true "user_id"
// @Success 200 {object} models.FavouriteModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/favourite/add [POST]
func (h HandlerV1) AddToFavourites(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "AddToFavourites")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	establishment_id := c.Query("establishment_id")
	user_id := c.Query("user_id")

	favourite_id := uuid.New().String()

	response, err := h.Service.EstablishmentService().AddToFavourites(ctx, &pb.AddToFavouritesRequest{
		Favourite: &pb.Favourite{
			FavouriteId:     favourite_id,
			EstablishmentId: establishment_id,
			UserId:          user_id,
		},
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	respModel := models.FavouriteModel{
		FavouriteId:     response.Favourite.FavouriteId,
		EstablishmentId: response.Favourite.EstablishmentId,
		UserId:          response.Favourite.UserId,
		CreatedAt:       response.Favourite.CreatedAt,
		UpdatedAt:       response.Favourite.UpdatedAt,
	}
	c.JSON(200, respModel)
}

// REMOVE FROM FAVOURITES BY FAVOURITE_ID
// @Summary REMOVE FROM FAVOURITES BY FAVOURITE_ID
// @Description Api for removing favourite by favourite_id
// @Tags FAVOURITE
// @Accept json
// @Produce json
// @Param favourite_id query string true "favourite_id"
// @Success 200 {object} models.RemoveResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/favourite/remove [DELETE]
func (h HandlerV1) RemoveFromFavourites(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "RemoveFromFavourites")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	favourite_id := c.Query("favourite_id")

	response, err := h.Service.EstablishmentService().RemoveFromFavourites(ctx, &pb.RemoveFromFavouritesRequest{
		FavouriteId: favourite_id,
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
			"message": "not removed",
		})
		h.Logger.Error("not removed")
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully removed",
	})
}

// LIST FAVOURITES BY USER_ID
// @Summary LIST FAVOURITES BY USER_ID
// @Description Api for listing favourites by favourite_id
// @Tags FAVOURITE
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Success 200 {object} models.ListFavouritesModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/favourite/list [GET]
func (h HandlerV1) ListFavouritesByUserId(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, span := otlp.Start(ctx, "api", "ListFavouritesByUserId")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
	)

	user_id := c.Query("user_id")

	response, err := h.Service.EstablishmentService().ListFavouritesByUserId(ctx, &pb.ListFavouritesByUserIdRequest{
		UserId: user_id,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	var favourites []*models.FavouriteModel

	for _, respFavourite := range response.Favourites {
		favourite := models.FavouriteModel{
			FavouriteId:     respFavourite.FavouriteId,
			EstablishmentId: respFavourite.EstablishmentId,
			UserId:          respFavourite.UserId,
			CreatedAt:       respFavourite.CreatedAt,
			UpdatedAt:       respFavourite.UpdatedAt,
		}

		favourites = append(favourites, &favourite)
	}

	c.JSON(200, favourites)
}
