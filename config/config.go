package config

import (
	"fmt"
	"os"
)

var (
	defaultHttpPort = "8080"
	HTTPDomain      = getEnvVar("HTTP_DOMAIN", fmt.Sprintf("http://localhost:%s", defaultHttpPort))
	HTTPPort        = getEnvVar("PORT", defaultHttpPort)
	GO_ENV          = getEnvVar("environment", "dev")
)

func getEnvVar(varName string, defaultValue string) string {
	value, ok := os.LookupEnv(varName)
	if !ok {
		if defaultValue == "" {
			panic(fmt.Sprintf("Cannot get %s env var", varName))
		}
		return defaultValue
	}

	return value

}
