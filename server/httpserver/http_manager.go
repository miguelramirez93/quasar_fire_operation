package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miguelramirez93/quasar_fire_operation/config"
	"github.com/miguelramirez93/quasar_fire_operation/handlers/message/delivery"
	_ "github.com/miguelramirez93/quasar_fire_operation/server/httpserver/docs"
	"github.com/miguelramirez93/quasar_fire_operation/shared/utils/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HTTPInitDocumentation(r *gin.Engine) {
	baseURL := config.HTTPDomain
	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", baseURL))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func HTTPInitHandlers(r *gin.Engine) {
	logger.Info("Initialising http controllers...")
	delivery.NewDecodeMessageAndSourceHttpController(r)
	delivery.NewAddSatelliteMessageHttpController(r)
	delivery.NewDecodeSplitedMessageAndSource(r)
	HTTPInitDocumentation(r)
}
