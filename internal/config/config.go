package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	pgDSN = "PG_DSN"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}
}

func GetPGDSN() string {
	return os.Getenv(pgDSN)
}
