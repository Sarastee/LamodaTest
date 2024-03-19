package config

import (
	"fmt"
	"net"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// GRPCConfigSearcher interface for search grpc config
type GRPCConfigSearcher interface {
	Get() (*GRPCConfig, error)
}

// LogConfigSearcher interface for search Log config.
type LogConfigSearcher interface {
	Get() (*LogConfig, error)
}

// PGConfigSearcher interface for search PG config.
type PGConfigSearcher interface {
	Get() (*PGConfig, error)
}

// Load dotenv from path to env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPCConfig grpc config.
type GRPCConfig struct {
	Host string
	Port string
}

// Address get address for grpc server.
func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

// LogConfig config for zerolog.
type LogConfig struct {
	LogLevel   zerolog.Level
	TimeFormat string
}

// PGConfig config for postgresql.
type PGConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// DSN ..
func (cfg *PGConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)
}
