package initializers

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var envVarsLoaded bool = false

func LoadEnvVars() {
	// Load environment variables

	if envVarsLoaded {
		return
	}

	envVarsLoaded = true

	if env := os.Getenv("ENV"); strings.ToLower(env) == "production" {
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
