package config

import (
	"flag"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
)

type (
	Config struct {
		HTTP
		Log
		DB
		Auth
	}

	HTTP struct {
		Port string
	}

	Log struct {
		Level string
	}

	DB struct {
		URL string
	}

	Auth struct {
		JWTSecret string
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	port := flag.String("port", "localhost:8080", "port")
	logLevel := flag.String("log_level", string(logger.InfoLevel), "logger level")
	dbURL := flag.String("db_url", "", "url for connection to database")

	// auth
	jwtSecret := flag.String("jwt_secret", "secret", "secret to sign jwt")

	flag.Parse()

	cfg := &Config{
		HTTP: HTTP{
			Port: *port,
		},
		Log: Log{
			Level: *logLevel,
		},
		DB: DB{
			URL: *dbURL,
		},
		Auth: Auth{
			JWTSecret: *jwtSecret,
		},
	}

	return cfg, nil
}
