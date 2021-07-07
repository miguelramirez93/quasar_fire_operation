package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miguelramirez93/quasar_fire_operation/shared/utils/logger"
)

// @title Test API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @BasePath /

// HTTPInitServer init http delivery server
func HTTPInitServer(port string) {
	logger.Info("starting http server", port)
	router := gin.Default()
	HTTPInitHandlers(router)

	router.Run(fmt.Sprintf(":%s", port))
	logger.Info("Server successfully started")
}
