package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/sarastee/LamodaTest/internal/config"
)

const (
	host     = "PG_HOST"
	port     = "PG_PORT"
	user     = "POSTGRES_USER"
	password = "POSTGRES_PASSWORD" // #nosec G101
	dbName   = "POSTGRES_DB"
)

type pgCfgSearcher struct{}

func NewPgCfgSearcher() *pgCfgSearcher {
	return &pgCfgSearcher{}
}

func (s *pgCfgSearcher) Get() (*config.PGConfig, error) {
	dbHost := os.Getenv(host)
	if len(dbHost) == 0 {
		return nil, errors.New("db host not found")
	}

	dbPort := os.Getenv(port)
	if len(dbHost) == 0 {
		return nil, errors.New("db port not found")
	}

	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, fmt.Errorf("incorrect port format: %v", err)
	}

	dbUser := os.Getenv(user)
	if len(dbUser) == 0 {
		return nil, errors.New("db user not found")
	}

	dbPass := os.Getenv(password)
	if len(dbPass) == 0 {
		return nil, errors.New("db password not found")
	}

	name := os.Getenv(dbName)
	if len(name) == 0 {
		return nil, errors.New("db name not found")
	}

	return &config.PGConfig{
		Host:     dbHost,
		Port:     dbPortInt,
		User:     dbUser,
		Password: dbPass,
		DbName:   name,
	}, nil
}
