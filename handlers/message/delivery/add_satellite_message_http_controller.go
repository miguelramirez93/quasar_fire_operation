package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	messagehandler "github.com/miguelramirez93/quasar_fire_operation/handlers/message"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

type AddSatelliteMessageHttpController struct {
	Handler messagehandler.AddSatelliteMessageHandler
}

func NewAddSatelliteMessageHttpController(r *gin.Engine) {
	controller := &AddSatelliteMessageHttpController{
		Handler: messagehandler.NewAddSatelliteMessageHandler(),
	}

	r.POST("/topsecret_split/:satellite_name", controller.AddSatelliteMessage)
}

// AddSatelliteMessage godoc
// @Summary adds satellite message handler
// @Description adds satellite message to data source, if there is a message to the same satellite it will be replaced with this one.
// @Accept  json
// @Produce  json
// @Param body body models.AddSatelliteMessageReq true "satellite message data"
// @Param satellite_name path string true "satellite name"
// @Success 200
// @Failure 403 {object} apperrors.DeliveryError
// @Failure 404 {object} apperrors.DeliveryError
// @Router /topsecret_split/:satellite_name [post]
func (ctr *AddSatelliteMessageHttpController) AddSatelliteMessage(c *gin.Context) {
	var reqBody models.AddSatelliteMessageReq
	satelliteName := c.Param("satellite_name")
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, apperrors.NewDeliveryError(apperrors.ErrBadInput))
		return
	}

	satelliteMessageData := addSatelliteReqToSatelliteMessage(satelliteName, reqBody)

	addedSatelliteData, err := ctr.Handler.AddSatelliteMessage(satelliteMessageData)

	if err != nil {
		c.JSON(http.StatusNotFound, apperrors.NewDeliveryError(err))
		return
	}

	c.JSON(http.StatusOK, addedSatelliteData)

}

func addSatelliteReqToSatelliteMessage(satelliteName string, addSatelliteReqData models.AddSatelliteMessageReq) models.SatelliteMessage {
	return models.SatelliteMessage{
		SatelliteName: satelliteName,
		Distance:      addSatelliteReqData.Distance,
		Message:       addSatelliteReqData.Message,
	}
}
