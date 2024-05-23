package v1

import (
	models "Booking/api-service-booking/api/models"
	pbu "Booking/api-service-booking/genproto/user-proto"
	"Booking/api-service-booking/internal/pkg/etc"
	l "Booking/api-service-booking/internal/pkg/logger"
	"Booking/api-service-booking/internal/pkg/otlp"
	tokens "Booking/api-service-booking/internal/pkg/token"
	"Booking/api-service-booking/internal/pkg/utils"
	valid "Booking/api-service-booking/internal/pkg/validation"

	// "context"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

// CREATE ADMIN
// @Summary CREATE ADMIN
// @Security BearerAuth
// @Description Api for Create admin
// @Tags ADMIN
// @Accept json
// @Produce json
// @Param Admin body models.UserReq true "createModel"
// @Success 200 {object} models.UserRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/admins [post]
func (h *HandlerV1) CreateAdmin(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "CreateAdmin")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
		attribute.Key("host").String(c.Request.Host),
	)
	defer span.End()

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

	res := valid.IsValidEmail(body.Email)
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Email. Try again",
		})

		h.Logger.Error("Incorrect Email. Try again, error while in Create")
		return
	}

	res = valid.IsValidPassword(body.Password)
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Password. Try again",
		})

		h.Logger.Error("Incorrect Password. Try again, error while in Create")
		return
	}

	isEmail, err := h.Service.UserService().CheckUniquess(ctx, &pbu.FV{
		Field: "email",
		Value: body.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Went wrong",
		})

		h.Logger.Error("Error while check unique email in Create")
		return
	}

	if isEmail.Code != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already in use",
		})

		return
	}

	password, err := etc.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Went wrong",
		})

		h.Logger.Error("Error while hash password in Create")
		return
	}

	newId := uuid.NewString()

	h.JwtHandler = tokens.JwtHandler{
		Sub:  newId,
		Iss:  "client",
		Role: "admin",
		SigninKey: h.Config.Token.SignInKey,
		Log:  h.Logger,
	}

	access, refresh, err := h.JwtHandler.GenerateJwt()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "error while generating jwt",
		})
		h.Logger.Error("error generate new jwt tokens", l.Error(err))
		return
	}

	response, err := h.Service.UserService().Create(ctx, &pbu.User{
		Id:           newId,
		FullName:     body.FullName,
		Email:        body.Email,
		Password:     password,
		DateOfBirth:  body.DateOfBirth,
		// ProfileImg:   body.ProfileImg,
		Card:         body.Card,
		Gender:       body.Gender,
		PhoneNumber:  body.PhoneNumber,
		Role:         "admin",
		RefreshToken: refresh,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	c.JSON(http.StatusCreated, &models.UserResCreate{
		Id:           response.Id,
		FullName:     response.FullName,
		Email:        response.Email,
		DateOfBirth:  response.DateOfBirth,
		ProfileImg:   response.ProfileImg,
		Card:         response.Card,
		Gender:       response.Gender,
		PhoneNumber:  response.PhoneNumber,
		Role:         response.Role,
		AccessToken:  access,
		RefreshToken: response.RefreshToken,
	})
}

// GET ADMIN
// @Summary GET ADMIN
// @Security BearerAuth
// @Description Api for Get
// @Tags ADMIN
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.UserRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/admins/{id} [get]
func (h *HandlerV1) GetAdmin(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "GetAdmin")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
		attribute.Key("host").String(c.Request.Host),
	)
	defer span.End()

	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")

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

	if response.User.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't get",
		})
		return
	}

	c.JSON(http.StatusOK, &models.UserRes{
		Id:           response.User.Id,
		FullName:     response.User.FullName,
		Email:        response.User.Email,
		DateOfBirth:  response.User.DateOfBirth,
		ProfileImg:   response.User.ProfileImg,
		Card:         response.User.Card,
		Gender:       response.User.Gender,
		PhoneNumber:  response.User.PhoneNumber,
		Role:         response.User.Role,
		RefreshToken: response.User.RefreshToken,
		CreatedAt:    response.User.CreatedAt,
		UpdatedAt:    response.User.CreatedAt,
		DeletedAt:    response.User.DeletedAt,
	})
}

// LIST ADMINS
// @Summary LIST ADMINS
// @Security BearerAuth
// @Description Api for ListAdmins
// @Tags ADMIN
// @Accept json
// @Produce json
// @Param request query models.Pagination true "request"
// @Success 200 {object} pbu.ListUsersRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/admins/list [get]
func (h *HandlerV1) ListAdmins(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "ListAdmins")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
		attribute.Key("host").String(c.Request.Host),
	)
	defer span.End()

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

	response, err := h.Service.UserService().ListUsers(
		ctx, &pbu.ListUsersReq{
			Limit:  params.Limit,
			Offset: (params.Page - 1) * params.Limit,
			Fv: &pbu.FV{
				Field: "role",
				Value: "admin",
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

// UPDATE ADMIN
// @Summary UPDATE ADMIN
// @Security BearerAuth
// @Description Api for Update
// @Tags ADMIN
// @Accept json
// @Produce json
// @Param Admin body models.UserReq true "createModel"
// @Success 200 {object} models.UserRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/admins [put]
func (h *HandlerV1) UpdateAdmin(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "UpdateAdmin")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
		attribute.Key("host").String(c.Request.Host),
	)
	defer span.End()

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
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

    userID, statusCode := GetIdFromToken(c.Request, h.Config)
	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{
            "error": "Can't get",
        })
        return
    }

	getUser, err := h.Service.UserService().Get(ctx, &pbu.Filter{
		Filter: map[string]string{"id": userID},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Went wrong",
		})
		h.Logger.Error("failed to get user in update", l.Error(err))
		return
	}

	if getUser.User.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't update",
		})
		return
	}

	if body.Password != "" {
		resPass := valid.IsValidPassword(body.Email)
		if !resPass {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Incorrect Password. Try again",
			})

			h.Logger.Error("Incorrect Password. Try again, error while in update")
			return
		}
		body.Password, err = etc.HashPassword(body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Went wrong",
			})
			h.Logger.Error("failed to hash password in update", l.Error(err))
			return
		}
	}

	if body.Email != "" {
		checkEmail, err := h.Service.UserService().CheckUniquess(ctx, &pbu.FV{
			Field:                "email",
			Value:                body.Email,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Went wrong",
			})
	
			h.Logger.Error("Error while check unique email in update admin")
			return
		}
	
		if checkEmail.Code != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email already in use",
			})
	
			return
		}

		res := valid.IsValidEmail(body.Email)
		if !res {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Incorrect Email. Try again",
			})

			h.Logger.Error("Incorrect Email. Try again, error while in update")
			return
		}
	}

	if body.FullName == "" {
		body.FullName = getUser.User.FullName
	}

	if body.Email == "" {
		body.Email = getUser.User.Email
	}

	if body.Password == "" {
		body.Password = getUser.User.Password
	}

	if body.DateOfBirth == "" {
		body.DateOfBirth = getUser.User.DateOfBirth
	}

	if body.Card == "" {
		body.Card = getUser.User.Card
	}

	if body.Gender == "" {
		body.Gender = getUser.User.Gender
	}

	if body.PhoneNumber == "" {
		body.PhoneNumber = getUser.User.PhoneNumber
	}

	response, err := h.Service.UserService().Update(ctx, &pbu.User{
		Id:          userID,
		FullName:    body.FullName,
		Email:       body.Email,
		Password:    body.Password,
		DateOfBirth: body.DateOfBirth,
		ProfileImg:  "",
		Card:        body.Card,
		Gender:      body.Gender,
		PhoneNumber: body.PhoneNumber,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, &models.UserRes{
		Id:           response.Id,
		FullName:     response.FullName,
		Email:        response.Email,
		DateOfBirth:  response.DateOfBirth,
		ProfileImg:   response.ProfileImg,
		Card:         response.Card,
		Gender:       response.Gender,
		PhoneNumber:  response.PhoneNumber,
		Role:         response.Role,
		RefreshToken: response.RefreshToken,
		CreatedAt:    response.CreatedAt,
		UpdatedAt:    response.UpdatedAt,
		DeletedAt:    response.DeletedAt,
	})
}

// DELETE ADMIN
// @Summary DELETE ADMIN
// @Security BearerAuth
// @Description Api for Delete
// @Tags ADMIN
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.RegisterRes
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/admins/{id} [delete]
func (h *HandlerV1) DeleteAdmin(c *gin.Context) {
	ctx, span := otlp.Start(c, "api", "DeleteAdmin")
	span.SetAttributes(
		attribute.Key("method").String(c.Request.Method),
		attribute.Key("host").String(c.Request.Host),
	)
	defer span.End()

	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")

	user, err := h.Service.UserService().Get(ctx, &pbu.Filter{
		Filter: map[string]string{"id": id},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Went wrong, error",
		})
		h.Logger.Error("failed to get admin in delete admin", l.Error(err))
		return
	}

	if user.User.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't delete",
		})
		return
	}

	_, err = h.Service.UserService().SoftDelete(
		ctx, &pbu.Id{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Went wrong, error",
		})
		h.Logger.Error("failed to delete user", l.Error(err))
		return
	}

	// if response != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Went wrong",
	// 	})
	// 	h.Logger.Error("failed to delete user", l.Error(err))
	// 	return
	// }

	c.JSON(http.StatusOK, &models.RegisterRes{
		Content: "User has been deleted",
	})
}
