package main

import (
	"github.com/miguelramirez93/quasar_fire_operation/config"
	"github.com/miguelramirez93/quasar_fire_operation/server/httpserver"
	"github.com/miguelramirez93/quasar_fire_operation/shared/utils/logger"
)

func main() {
	logger.Init(logger.LogLevel(config.GO_ENV))
	logger.Info("Starting app...")
	httpserver.HTTPInitServer(config.HTTPPort)
}
