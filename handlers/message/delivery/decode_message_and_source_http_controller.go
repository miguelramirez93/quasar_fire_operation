package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	messagehandler "github.com/miguelramirez93/quasar_fire_operation/handlers/message"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

type DecodeMessageAndSourceHttpController struct {
	Handler messagehandler.DecodeMessageAndSourceHandler
}

func NewDecodeMessageAndSourceHttpController(r *gin.Engine) {
	controller := &DecodeMessageAndSourceHttpController{
		Handler: messagehandler.NewDecodeMessageAndSourceHandler(),
	}

	r.POST("/topsecret", controller.DecodeMessageAndSource)
}

// DecodeMessageAndSource godoc
// @Summary decode message from multiple sources
// @Description return the message and coordenates for emisor
// @Accept  json
// @Produce  json
// @Param body body models.DecodeMsgAndSourceReq true "satellites message data"
// @Success 200
// @Failure 403 {object} apperrors.DeliveryError
// @Failure 404 {object} apperrors.DeliveryError
// @Router /topsecret [post]
func (ctr *DecodeMessageAndSourceHttpController) DecodeMessageAndSource(c *gin.Context) {
	var reqBody models.DecodeMsgAndSourceReq

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, apperrors.NewDeliveryError(apperrors.ErrBadInput))
		return
	}

	decodedMsg, err := ctr.Handler.DecodeMessageAndSource(reqBody.Satellites)

	if err != nil {
		c.JSON(http.StatusNotFound, apperrors.NewDeliveryError(err))
		return
	}

	c.JSON(http.StatusOK, decodedMsg)

}
