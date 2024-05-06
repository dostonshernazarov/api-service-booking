package v1

// import (
// 	"context"
// 	"strconv"
// 	"time"

// 	"Booking/api-service-booking/api/models"
// 	pbc "Booking/api-service-booking/genproto/client_service"
// 	pbj "Booking/api-service-booking/genproto/jobs_service"
// 	"Booking/api-service-booking/internal/pkg/otlp"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"go.opentelemetry.io/otel/attribute"
// 	"google.golang.org/protobuf/encoding/protojson"
// )

// // CREATE CLIENT
// // @Summary CREATE CLIENT
// // @Description Api for creating client
// // @Tags CLIENT
// // @Accept json
// // @Produce json
// // @Param Client body models.CreateClient true "Client"
// // @Success 200 {object} models.ClientModel
// // @Failure 404 {object} models.StandartError
// // @Failure 500 {object} models.StandartError
// // @Router /v1/client/create [POST]
// func (h HandlerV1) Create(c *gin.Context) {
// 	var (
// 		clientBody  models.CreateClient
// 		jspbMarshal protojson.MarshalOptions
// 	)

// 	jspbMarshal.UseProtoNames = true

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	ctx, span := otlp.Start(ctx, "api", "CreateClient")
// 	span.SetAttributes(
// 		attribute.Key("method").String(c.Request.Method),
// 	)

// 	if err := c.ShouldBindJSON(&clientBody); err != nil {
// 		c.JSON(404, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	client_id := uuid.New().String()

// 	respClient, err := h.Service.ClientService().Create(ctx, &pbc.Client{
// 		ClientId:    client_id,
// 		FullName:    clientBody.FullName,
// 		Username:    clientBody.Username,
// 		Email:       clientBody.Email,
// 		Password:    clientBody.Password,
// 		DateOfBirth: clientBody.DateOfBirth.Format("2006-01-02T15:04:05Z"),
// 		Address:     clientBody.Address,
// 		ImageUrl:    clientBody.ImageUrl,
// 	})
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	respModel := models.ClientModel{
// 		ClientId:    respClient.ClientId,
// 		FullName:    respClient.FullName,
// 		Username:    respClient.Username,
// 		Email:       respClient.Email,
// 		Password:    respClient.Password,
// 		DateOfBirth: respClient.DateOfBirth,
// 		Address:     respClient.Address,
// 		ImageUrl:    respClient.ImageUrl,
// 		CreatedAt:   respClient.CreatedAt,
// 		UpdatedAt:   respClient.UpdatedAt,
// 	}

// 	c.JSON(200, respModel)
// }

// // GET CLIENT BY CLIENT_ID
// // @Summary GET CLIENT BY CLIENT_ID
// // @Description Api for getting client by client_id
// // @Tags CLIENT
// // @Accept json
// // @Produce json
// // @Param client_id query string true "client_id"
// // @Success 200 {object} models.GetClientModel
// // @Failure 404 {object} models.StandartError
// // @Failure 500 {object} models.StandartError
// // @Router /v1/client/get [GET]
// func (h HandlerV1) Get(c *gin.Context) {
// 	var (
// 		jspbMarshal protojson.MarshalOptions
// 	)

// 	jspbMarshal.UseProtoNames = true

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	ctx, span := otlp.Start(ctx, "api", "GetClient")
// 	span.SetAttributes(
// 		attribute.Key("method").String(c.Request.Method),
// 	)

// 	client_id := c.Query("client_id")

// 	response, err := h.Service.ClientService().Get(ctx, &pbc.GetRequest{
// 		ClientId: client_id,
// 	})
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.Logger.Error(err.Error())
// 		return
// 	}

// 	respJobs, err := h.Service.JobsService().ListByClientId(ctx, &pbj.ListByClientIdRequest{
// 		ClientId: response.Client.ClientId,
// 	})
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.Logger.Error(err.Error())
// 		return
// 	}

// 	var jobs []*models.JobsModel

// 	for _, respJob := range respJobs.Jobs {
// 		job := models.JobsModel{
// 			JobId:        respJob.JobId,
// 			ClientId:     respJob.ClientId,
// 			Company:      respJob.Company,
// 			Position:     respJob.Position,
// 			TimeInfo:     respJob.TimeInfo,
// 			IsWorkingNow: respJob.IsWorkingNow,
// 			CreatedAt:    respJob.CreatedAt,
// 			UpdatedAt:    respJob.UpdatedAt,
// 		}

// 		jobs = append(jobs, &job)
// 	}

// 	respModel := models.GetClientModel{
// 		ClientId:    response.Client.ClientId,
// 		FullName:    response.Client.FullName,
// 		Username:    response.Client.Username,
// 		Email:       response.Client.Email,
// 		Password:    response.Client.Password,
// 		DateOfBirth: response.Client.DateOfBirth,
// 		Address:     response.Client.Address,
// 		ImageUrl:    response.Client.ImageUrl,
// 		Jobs:        jobs,
// 	}

// 	c.JSON(200, respModel)
// }

// // LIST CLIENT BY PAGE AND LIMIT
// // @Summary LIST CLIENT BY CLIENT_ID
// // @Description Api for listing client by page and limit
// // @Tags CLIENT
// // @Accept json
// // @Produce json
// // @Param page query string true "page"
// // @Param limit query string true "limit"
// // @Success 200 {object} models.ListModel
// // @Failure 404 {object} models.StandartError
// // @Failure 500 {object} models.StandartError
// // @Router /v1/client/list [GET]
// func (h HandlerV1) List(c *gin.Context) {
// 	var (
// 		jspbMarshal protojson.MarshalOptions
// 	)

// 	jspbMarshal.UseProtoNames = true

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	ctx, span := otlp.Start(ctx, "api", "ListClients")
// 	span.SetAttributes(
// 		attribute.Key("method").String(c.Request.Method),
// 	)

// 	page := c.Query("page")
// 	pageInt, err := strconv.Atoi(page)
// 	if err != nil {
// 		c.JSON(404, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.Logger.Error(err.Error())
// 		return
// 	}

// 	limit := c.Query("limit")
// 	limitInt, err := strconv.Atoi(limit)
// 	if err != nil {
// 		c.JSON(404, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.Logger.Error(err.Error())
// 		return
// 	}

// 	offset := (pageInt - 1) * limitInt

// 	response, err := h.Service.ClientService().List(ctx, &pbc.ListRequest{
// 		Offset: int64(offset),
// 		Limit:  int64(limitInt),
// 	})
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.Logger.Error(err.Error())
// 		return
// 	}

// 	var clients []*models.GetClientModel

// 	for _, respClient := range response.Clients {

// 		client := &models.GetClientModel{
// 			ClientId:    respClient.ClientId,
// 			FullName:    respClient.FullName,
// 			Username:    respClient.Username,
// 			Email:       respClient.Email,
// 			Password:    respClient.Password,
// 			DateOfBirth: respClient.DateOfBirth,
// 			Address:     respClient.Address,
// 			ImageUrl:    respClient.ImageUrl,
// 			CreatedAt:   respClient.CreatedAt,
// 			UpdatedAt:   respClient.UpdatedAt,
// 		}

// 		respJobs, err := h.Service.JobsService().ListByClientId(ctx, &pbj.ListByClientIdRequest{
// 			ClientId: respClient.ClientId,
// 		})
// 		if err != nil {
// 			c.JSON(404, gin.H{
// 				"error": err.Error(),
// 			})
// 			h.Logger.Error(err.Error())
// 			return
// 		}

// 		var jobs []*models.JobsModel

// 		for _, respJob := range respJobs.Jobs {
// 			job := &models.JobsModel{
// 				JobId:        respJob.JobId,
// 				ClientId:     respJob.ClientId,
// 				Company:      respJob.Company,
// 				Position:     respJob.Position,
// 				TimeInfo:     respJob.TimeInfo,
// 				IsWorkingNow: respJob.IsWorkingNow,
// 				CreatedAt:    respJob.CreatedAt,
// 				UpdatedAt:    respJob.UpdatedAt,
// 			}

// 			jobs = append(jobs, job)
// 		}

// 		client.Jobs = jobs

// 		clients = append(clients, client)
// 	}

// 	c.JSON(200, clients)
// }

// // UPDATE CLIENT
// // @Summary UPDATE CLIENT
// // @Description Api for updating client by client_id
// // @Tags CLIENT
// // @Accept json
// // @Produce json
// // @Param UpdatingClient body models.CreateClient true "UpdatingClient"
// // @Param client_id query string true "client_id"
// // @Success 200 {object} models.ClientModel
// // @Failure 404 {object} models.StandartError
// // @Failure 500 {object} models.StandartError
// // @Router /v1/client/update [PATCH]
// func (h HandlerV1) Update(c *gin.Context) {
// 	var (
// 		clientBody  models.CreateClient
// 		jspbMarshal protojson.MarshalOptions
// 	)

// 	jspbMarshal.UseProtoNames = true

// 	client_id := c.Query("client_id")

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	ctx, span := otlp.Start(ctx, "api", "UpdateClient")
// 	span.SetAttributes(
// 		attribute.Key("method").String(c.Request.Method),
// 	)

// 	err := c.ShouldBindJSON(&clientBody)

// 	if err != nil {
// 		c.JSON(404, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	response, err := h.Service.ClientService().Update(ctx, &pbc.UpdateRequest{
// 		Client: &pbc.Client{
// 			ClientId:    client_id,
// 			FullName:    clientBody.FullName,
// 			Username:    clientBody.Username,
// 			Email:       clientBody.Email,
// 			Password:    clientBody.Password,
// 			DateOfBirth: clientBody.DateOfBirth.String(),
// 			Address:     clientBody.Address,
// 			ImageUrl:    clientBody.ImageUrl,
// 		},
// 	})
// 	if err != nil {
// 		c.JSON(200, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.Logger.Error(err.Error())
// 		return
// 	}

// 	respModel := models.ClientModel{
// 		ClientId:    response.Client.ClientId,
// 		FullName:    response.Client.FullName,
// 		Username:    response.Client.Username,
// 		Email:       response.Client.Email,
// 		Password:    response.Client.Password,
// 		DateOfBirth: response.Client.DateOfBirth,
// 		Address:     response.Client.Address,
// 		ImageUrl:    response.Client.ImageUrl,
// 		CreatedAt:   response.Client.CreatedAt,
// 		UpdatedAt:   response.Client.UpdatedAt,
// 	}

// 	c.JSON(200, respModel)
// }

// // DELETE CLIENT BY CLIENT_ID
// // @Summary DELETE CLIENT BY CLIENT_ID
// // @Description Api for deleting client by client_id
// // @Tags CLIENT
// // @Accept json
// // @Produce json
// // @Param client_id query string true "client_id"
// // @Success 200 {object} models.DeleteResponse
// // @Failure 404 {object} models.StandartError
// // @Failure 500 {object} models.StandartError
// // @Router /v1/client/delete [DELETE]
// func (h HandlerV1) Delete(c *gin.Context) {
// 	var (
// 		jspbMarshal protojson.MarshalOptions
// 	)

// 	jspbMarshal.UseProtoNames = true

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	ctx, span := otlp.Start(ctx, "api", "DeleteClient")
// 	span.SetAttributes(
// 		attribute.Key("method").String(c.Request.Method),
// 	)

// 	client_id := c.Query("client_id")

// 	// firstly, we need to delete client's job
// 	// so, here I initially received all jobs by client_id
// 	rJobs, _ := h.Service.JobsService().ListByClientId(ctx, &pbj.ListByClientIdRequest{
// 		ClientId: client_id,
// 	})

// 	for _, rJob := range rJobs.Jobs {
// 		_, _ = h.Service.JobsService().Delete(ctx, &pbj.DeleteRequest{
// 			JobId: rJob.JobId,
// 		})
// 	}

// 	response, err := h.Service.ClientService().Delete(ctx, &pbc.DeleteRequest{
// 		ClientId: client_id,
// 	})
// 	if err != nil {
// 		if err.Error() == "rpc error: code = Unknown desc = no sql rows" {
// 			c.JSON(404, gin.H{
// 				"error": "this client has already been deleted",
// 			})
// 			h.Logger.Error(err.Error())
// 			return
// 		} else {
// 			c.JSON(500, gin.H{
// 				"error": err.Error(),
// 			})
// 			h.Logger.Error(err.Error())
// 			return
// 		}
// 	}

// 	if !response.Success {
// 		c.JSON(404, gin.H{
// 			"error": "not deleted",
// 		})
// 		h.Logger.Error("not deleted")
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "client successfuly deleted",
// 	})
// }
