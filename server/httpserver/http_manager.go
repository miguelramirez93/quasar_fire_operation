package httpserver

import (
	"github.com/gin-gonic/gin"
	messagehandler "github.com/miguelramirez93/quasar_fire_operation/handlers/message"
	"github.com/miguelramirez93/quasar_fire_operation/handlers/message/delivery"
	repositories "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite"
	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
)

func HTTPInitHandlers(r *gin.Engine) {
	//repositories
	satelliteRepository := repositories.NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
	//decode_message_and_source
	decodeMessageAndSourceHandler := messagehandler.NewDecodeMessageAndSourceHandler(satelliteRepository)

	delivery.NewDecodeMessageAndSourceHttpController(r, decodeMessageAndSourceHandler)

}
