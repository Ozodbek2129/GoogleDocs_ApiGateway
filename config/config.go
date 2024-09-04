package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE   string
	DOCS_SERVICE string
	SIGNING_KEY    string
	API_GATEWAY    string
}

func Load() Config {	
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":1234"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "nimadurGo11"))
	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", ":9876"))
	config.DOCS_SERVICE = cast.ToString(Coalesce("DOCS_SERVICE", ":50052"))
	

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}