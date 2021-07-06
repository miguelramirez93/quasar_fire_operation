package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

	serverPort := "8080"

	router.Run(fmt.Sprintf(":%s", serverPort))
}
