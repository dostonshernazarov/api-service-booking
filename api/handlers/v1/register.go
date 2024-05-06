package v1

import "github.com/gin-gonic/gin"

// REGISTER USER
// @Summary REGISTER USER
// @Description Api for register new user
// @Tags SINUP
// @Accept json
// @Produce json
// @Param Attraction body models.CreateAttraction true "Attraction"
// @Param owner_id query string true "owner_id"
// @Success 200 {object} models.AttractionModel
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/attraction/create [POST]
func (h HandlerV1) RegisterUser(c *gin.Context) {
	

	c.JSON(200, nil)
}