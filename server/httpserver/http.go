package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miguelramirez93/quasar_fire_operation/config"
)

// @title Test API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @BasePath /

// HTTPInitServer init http delivery server
func HTTPInitServer() {
	router := gin.Default()

	HTTPInitHandlers(router)

	serverPort := config.HTTPPort

	router.Run(fmt.Sprintf(":%s", serverPort))
}
