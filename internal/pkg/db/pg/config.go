package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

var pgConnStr = goDotEnvVar("POSTGRES_URL")
