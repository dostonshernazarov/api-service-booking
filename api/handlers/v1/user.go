package v1
<<<<<<< HEAD

import (
	models "Booking/api-service-booking/api/models"
	pbu "Booking/api-service-booking/genproto/user-proto"
	l "Booking/api-service-booking/internal/pkg/logger"
	"Booking/api-service-booking/internal/pkg/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create
// @Summary Create
// @Security ApiKeyAuth
// @Description Api for Create
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.UserReq true "createModel"
// @Success 200 {object} models.UserRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/user/create [post]
func (h *HandlerV1) Create(c *gin.Context) {
	var (
		body        models.UserReq
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.UserService().Create(ctx, &pbu.User{
		FullName:		body.FullName,
		Email:			body.Email,
        Password:		body.Password,
        DateOfBirth:	body.DateOfBirth,
        ProfileImg:		body.ProfileImg,
        Card:			body.Card,
		Gender:			body.Gender,
        PhoneNumber:	body.PhoneNumber,
	})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusCreated, response)

	var userBody models.UserRes
	c.JSON(http.StatusCreated, &models.UserRes{
		Id:				userBody.Id,
		FullName:		userBody.FullName,
		Email:			userBody.Email,
        Password:		userBody.Password,
        DateOfBirth:	userBody.DateOfBirth,
        ProfileImg:		userBody.ProfileImg,
        Card:			userBody.Card,
		Gender:			userBody.Gender,
        PhoneNumber:	userBody.PhoneNumber,
        Role:			userBody.Role,
        RefreshToken:	userBody.RefreshToken,
        CreatedAt:		userBody.CreatedAt,
	})
}

// Get
// @Summary Get
// @Security ApiKeyAuth
// @Description Api for Get
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.UserRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/user/{id} [get]
func (h *HandlerV1) Get(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.UserService().Get(
		ctx, &pbu.Filter{
			Filter: map[string]string{"id": id},
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

// ListUsers
// @Summary ListUsers
// @Security ApiKeyAuth
// @Description Api for ListUsers
// @Tags user
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.Users
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/user/list/users [get]
func (h *HandlerV1) ListUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParam(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		return
	}
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.UserService().ListUsers(
		ctx, &pbu.ListUsersReq{
			Limit: params.Limit,
			Offset:  params.Offset,
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

// ListDeletedUsers
// @Summary ListDeletedUsers
// @Security ApiKeyAuth
// @Description Api for ListDeletedUsers
// @Tags user
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} models.Users
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/user/list/deleted [get]
func (h *HandlerV1) ListDeletedUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParam(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		return
	}
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.UserService().ListDeletedUsers(
		ctx, &pbu.ListUsersReq{
			Limit: params.Limit,
			Offset:  params.Offset,
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

// Update
// @Summary Update
// @Security ApiKeyAuth
// @Description Api for Update
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.UserReq true "createModel"
// @Success 200 {object} models.UserRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/user/update [put]
func (h *HandlerV1) Update(c *gin.Context) {
	var (
		body        pbu.User
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		// h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.UserService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete
// @Summary Delete
// @Security ApiKeyAuth
// @Description Api for Delete
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.StandartError
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/user/delete/{id} [delete]
func (h *HandlerV1) Delete(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*time.Duration(h.ContextTimeout))
	defer cancel()

	response, err := h.Service.UserService().SoftDelete(
		ctx, &pbu.Id{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
=======
>>>>>>> main
