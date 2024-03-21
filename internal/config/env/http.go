package env

import (
	"errors"
	"os"

	"github.com/sarastee/LamodaTest/internal/config"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

// HTTPCfgSearcher searcher for HTTP config.
type HTTPCfgSearcher struct{}

// NewHTTPCfgSearcher get instance for HTTP config searcher.
func NewHTTPCfgSearcher() *HTTPCfgSearcher {
	return &HTTPCfgSearcher{}
}

// Get searcher for grpc config.
func (s *HTTPCfgSearcher) Get() (*config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &config.HTTPConfig{
		Host: host,
		Port: port,
	}, nil
}
