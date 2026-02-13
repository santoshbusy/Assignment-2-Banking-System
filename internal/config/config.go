package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDsn string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	dsn := "host=" + os.Getenv("DB_HOST") + " " +
		"user=" + os.Getenv("DB_USER") + " " +
		"password=" + os.Getenv("DB_PASSWORD") + " " +
		"dbname=" + os.Getenv("DB_NAME") + " " +
		"port=" + os.Getenv("DB_PORT") + " " +
		"sslmode=" + os.Getenv("DB_SSLMODE")

	return Config{
		DBDsn: dsn,
	}
}
