// Package config provides types and functionalities for getting server configuration values from environment variables
package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"time"
)

// DatabaseConfig represents database configuration values
type DatabaseConfig struct {
	Dsn string `env:"DATABASE_DSN,default=host=localhost user=sajjad password=sajjad123 dbname=omp port=5432 sslmode=require TimeZone=UTC"`
}

// HTTPServerConfig represents HTTP server configuration values
type HTTPServerConfig struct {
	ReadTimeout       time.Duration `env:"HTTP_SERVER_READ_TIMEOUT,default=15s"`
	ReadHeaderTimeout time.Duration `env:"HTTP_SERVER_READ_HEADER_TIMEOUT,default=15s"`
	WriteTimeout      time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT,default=15s"`
	Address           string        `env:"HTTP_SERVER_ADDRESS,default=127.0.0.1"`
	Port              uint          `env:"HTTP_SERVER_PORT,default=4444"`
}
type Config struct {
	Database   DatabaseConfig
	HTTPServer HTTPServerConfig
}

// NewConfigFromEnv returns a *Config set by environment variables
func NewConfigFromEnv(ctx context.Context) (Config, error) {
	var c Config
	err := envconfig.Process(ctx, &c)
	return c, err
}
