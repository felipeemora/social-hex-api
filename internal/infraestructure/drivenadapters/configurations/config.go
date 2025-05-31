package configurations

import (
	"log"

	"github.com/joho/godotenv"
)

type Configurations struct {
	ServerPort string
	DBConfig   dbConfig
}

type dbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func Load() *Configurations {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configurations := &Configurations{
		ServerPort: GetString("SERVER_PORT", ":8080"),
		DBConfig: dbConfig{
			Addr:         GetString("DB_ADDR", ""),
			MaxOpenConns: GetInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns: GetInt("DB_MAX_IDLE_CONNS", 10),
			MaxIdleTime:  GetString("DB_MAX_IDLE_TIME", "5m"),
		},
	}

	return configurations
}
