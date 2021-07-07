package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	messagehandler "github.com/miguelramirez93/quasar_fire_operation/handlers/message"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
)

type DecodeSplitedMessageAndSourceHttpController struct {
	Handler messagehandler.DecodeSplitedMessageAndSourceHandler
}

func NewDecodeSplitedMessageAndSource(r *gin.Engine) {
	controller := &DecodeSplitedMessageAndSourceHttpController{
		Handler: messagehandler.NewDecodeSplitMessageAndSourceHandler(),
	}

	r.GET("/topsecret_split", controller.DecodeSplitedMessageAndSource)
}

// DecodeSplitedMessageAndSource godoc
// @Summary decode message from multiple sources (splited)
// @Description return the message and coordenates for emisor if there is stored messages previusly
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 403 {object} apperrors.DeliveryError
// @Failure 404 {object} apperrors.DeliveryError
// @Router /topsecret_split [get]
func (ctr *DecodeSplitedMessageAndSourceHttpController) DecodeSplitedMessageAndSource(c *gin.Context) {

	decodedMsg, err := ctr.Handler.DecodeSplitedMessageAndSource()

	if err != nil {
		c.JSON(http.StatusNotFound, apperrors.NewDeliveryError(err))
		return
	}

	c.JSON(http.StatusOK, decodedMsg)

}
