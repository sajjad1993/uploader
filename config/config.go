// Package config provides types and functionalities for getting server configuration values from environment variables
package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

// AlbumAPIConfig represents iTunes API configuration values
type AlbumAPIConfig struct {
	Address         string        `env:"ALBUM_API_ADDRESS,required"`
	SearchMethod    string        `env:"ALBUM_API_SEARCH_METHOD,default=search"`
	MediaTypeParam  string        `env:"ALBUM_API_MEDIA_TYPE_PARAM,default=media"`
	MediaType       string        `env:"ALBUM_API_MEDIA_TYPE,default=music"`
	EntityTypeParam string        `env:"ALBUM_API_ENTITY_TYPE_PARAM,default=entity"`
	EntityType      string        `env:"ALBUM_API_ENTITY_TYPE,default=album"`
	AttributeParam  string        `env:"ALBUM_API_ATTRIBUTE_PARAM,default=attribute"`
	Attribute       string        `env:"ALBUM_API_ATTRIBUTE,default=albumTerm"`
	SearchParam     string        `env:"ALBUM_API_SEARCH_PARAM,default=term"`
	MaxResultsParam string        `env:"ALBUM_API_MAX_RESULTS_PARAM,default=limit"`
	MaxResults      uint16        `env:"ALBUM_API_MAX_RESULTS,default=5"`
	CheckTimeout    time.Duration `env:"ALBUM_API_CHECK_TIMEOUT,default=10s"`
	CacheDuration   time.Duration `env:"ALBUM_API_CACHE_DURATION,default=30m"`
}

// BookAPIConfig represents Google books API configuration values
type BookAPIConfig struct {
	Address          string        `env:"BOOK_API_ADDRESS,required"`
	Version          string        `env:"BOOK_API_VERSION,default=v1"`
	SearchMethod     string        `env:"BOOK_API_SEARCH_METHOD,default=volumes"`
	PrintTypeParam   string        `env:"BOOK_API_PRINT_TYPE_PARAM,default=printType"`
	PrintType        string        `env:"BOOK_API_PRINT_TYPE,default=books"`
	FilterParam      string        `env:"BOOK_API_FILTER_PARAM,default=filter"`
	Filter           string        `env:"BOOK_API_FILTER,default=full"`
	SearchParam      string        `env:"BOOK_API_SEARCH_PARAM,default=q"`
	TitleSearchParam string        `env:"BOOK_API_TITLE_SEARCH_PARAM,default=intitle"`
	MaxResultsParam  string        `env:"BOOK_API_MAX_RESULTS_PARAM,default=maxResults"`
	MaxResults       uint16        `env:"BOOK_API_MAX_RESULTS,default=5"`
	CheckTimeout     time.Duration `env:"BOOK_API_CHECK_TIMEOUT,default=10s"`
	CacheDuration    time.Duration `env:"BOOK_API_CACHE_DURATION,default=30m"`
}

// HTTPClientConfig represents HTTP client configuration values
type HTTPClientConfig struct {
	MaxIdleConns        int           `env:"MAX_HTTP_IDLE_CONNECTIONS,default=100"`
	MaxConnsPerHost     int           `env:"MAX_HTTP_CONNECTIONS_PER_HOST,default=100"`
	MaxIdleConnsPerHost int           `env:"MAX_HTTP_IDLE_CONNECTIONS_PER_HOST,default=100"`
	Timeout             time.Duration `env:"HTTP_CLIENT_TIMEOUT,default=10s"`
}

// HTTPServerConfig represents HTTP server configuration values
type HTTPServerConfig struct {
	ReadTimeout       time.Duration `env:"HTTP_SERVER_READ_TIMEOUT,default=15s"`
	ReadHeaderTimeout time.Duration `env:"HTTP_SERVER_READ_HEADER_TIMEOUT,default=15s"`
	WriteTimeout      time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT,default=15s"`
	Address           string        `env:"HTTP_SERVER_ADDRESS,default=0.0.0.0"`
	Port              string        `env:"HTTP_SERVER_PORT,default=8080"`
}

// LogConfig represents log configuration values
type LogConfig struct {
	StackTraceLevel int8 `env:"LOG_STACK_TRACE_LEVEL,default=2"` //default is error level
	Level           int8 `env:"LOG_LEVEL,default=-1"`            //default is debug level
}

// MetricConfig represents metric configuration values
type MetricConfig struct {
	Namespace string `env:"METRIC_NAMESPACE,default=Kramp_Hub"`
	Subsystem string `env:"METRIC_SUBSYSTEM,default=content_service"`
	Name      string `env:"METRIC_NAME,default=content_api"`
}

// RedisConfig represents redis configuration values
type RedisConfig struct {
	Network     string        `env:"REDIS_NETWORK,default=tcp"`
	Address     string        `env:"REDIS_ADDRESS,default=127.0.0.1:6379"`
	DB          int           `env:"REDIS_DB,default=1"`
	Username    string        `env:"REDIS_USERNAME"`
	Password    string        `env:"REDIS_PASSWORD"`
	PingTimeout time.Duration `env:"REDIS_PING_TIMEOUT,default=10s"`
	MaxRetries  int           `env:"REDIS_MAX_RETRIES,default=3"`
	PoolSize    int           `env:"REDIS_POOL_SIZE,default=10"`
}

// Config represents server configuration values
type Config struct {
	AlbumAPI   AlbumAPIConfig
	BookAPI    BookAPIConfig
	HTTPClient HTTPClientConfig
	HTTPServer HTTPServerConfig
	Log        LogConfig
	Metric     MetricConfig
	Redis      RedisConfig
}

// NewConfigFromEnv returns a *Config set by environment variables
func NewConfigFromEnv(ctx context.Context) (Config, error) {
	var c Config
	err := envconfig.Process(ctx, &c)
	return c, err
}
