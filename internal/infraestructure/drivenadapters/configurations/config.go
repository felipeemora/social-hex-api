package configurations

import (
	"log"

	"github.com/joho/godotenv"
)

type Configurations struct {
	ServerPort string
	ApiURL     string
	DBConfig   *dbConfig
	MailConfig *MailConfig
	JWTConfig  *JWTConfig
}

type dbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type MailConfig struct {
	FromEmail string
	ApiKey    string
	Host      string
	Port      int
	Username  string
}

type JWTConfig struct {
	SecretKey string
	Issuer    string
	Audience  string
}

func Load() *Configurations {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configurations := &Configurations{
		ServerPort: GetString("SERVER_PORT", ":8080"),
		ApiURL:     GetString("API_URL", "localhost:8080"),
		DBConfig: &dbConfig{
			Addr:         GetString("DB_ADDR", ""),
			MaxOpenConns: GetInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns: GetInt("DB_MAX_IDLE_CONNS", 10),
			MaxIdleTime:  GetString("DB_MAX_IDLE_TIME", "5m"),
		},
		MailConfig: &MailConfig{
			FromEmail: GetString("MAIL_FROM_EMAIL", ""),
			ApiKey:    GetString("MAIL_API_KEY", ""),
			Host:      GetString("MAIL_HOST", ""),
			Port:      GetInt("MAIL_PORT", 587),
			Username:  GetString("MAIL_USERNAME", ""),
		},
		JWTConfig: &JWTConfig{
			SecretKey: GetString("JWT_SECRET_KEY", ""),
			Issuer:    GetString("JWT_ISSUER", ""),
			Audience:  GetString("JWT_AUDIENCE", ""),
		},
	}

	return configurations
}
