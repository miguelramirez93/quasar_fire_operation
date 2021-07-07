package httpserver

import (
	"github.com/gin-gonic/gin"
	messagehandler "github.com/miguelramirez93/quasar_fire_operation/handlers/message"
	"github.com/miguelramirez93/quasar_fire_operation/handlers/message/delivery"
)

func HTTPInitHandlers(r *gin.Engine) {
	//decode_message_and_source
	decodeMessageAndSourceHandler := messagehandler.NewDecodeMessageAndSourceHandler()

	addSatelliteMessageHandler := messagehandler.NewAddSatelliteMessageHandler()

	delivery.NewDecodeMessageAndSourceHttpController(r, decodeMessageAndSourceHandler)
	delivery.NewAddSatelliteMessageHttpController(r, addSatelliteMessageHandler)

}
