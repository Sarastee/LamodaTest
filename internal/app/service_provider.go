package app

import (
	"context"
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/sarastee/LamodaTest/internal/api/warehouse"
	"github.com/sarastee/LamodaTest/internal/config"
	"github.com/sarastee/LamodaTest/internal/config/env"
	"github.com/sarastee/LamodaTest/internal/repository"
	reserveRepository "github.com/sarastee/LamodaTest/internal/repository/reserve"
	warehouseRepository "github.com/sarastee/LamodaTest/internal/repository/warehouse"
	"github.com/sarastee/LamodaTest/internal/service"
	warehouseService "github.com/sarastee/LamodaTest/internal/service/warehouse"
	"github.com/sarastee/platform_common/pkg/closer"
	"github.com/sarastee/platform_common/pkg/db"
	"github.com/sarastee/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	pgConfig   *config.PgConfig
	grpcConfig *config.GRPCConfig
	logger     *zerolog.Logger

	dbClient  db.Client
	txManager db.TxManager

	reserveRepo   repository.ReserveRepository
	warehouseRepo repository.WarehouseRepository

	warehouseService service.WarehouseService

	warehouseImpl *warehouse.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PgConfig ..
func (s *serviceProvider) PgConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get PG config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig ..
func (s *serviceProvider) GRPCConfig() *config.GRPCConfig {
	if s.grpcConfig == nil {
		cfgSearcher := env.NewGRPCCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get gRPC config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// Logger ..
func (s *serviceProvider) Logger() *zerolog.Logger {
	if s.logger == nil {
		cfgSearcher := env.NewLogCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Logger config: %s", err.Error())
		}

		s.logger = setupZeroLog(cfg)
	}

	return s.logger
}

// DBClient ..
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PgConfig().DSN(), s.Logger())
		if err != nil {
			log.Fatalf("failure while creating DB: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("no connection to DB: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager ..
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = pg.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) WarehouseRepository(ctx context.Context) repository.WarehouseRepository {
	if s.warehouseRepo == nil {
		s.warehouseRepo = warehouseRepository.NewRepo(s.Logger(), s.DBClient(ctx))
	}

	return s.warehouseRepo
}

func (s *serviceProvider) ReserveRepository(ctx context.Context) repository.ReserveRepository {
	if s.reserveRepo == nil {
		s.reserveRepo = reserveRepository.NewRepo(s.Logger(), s.DBClient(ctx))
	}

	return s.reserveRepo
}

func (s *serviceProvider) WarehouseService(ctx context.Context) service.WarehouseService {
	if s.warehouseService == nil {
		s.warehouseService = warehouseService.NewService(
			s.Logger(),
			s.DBClient(ctx),
			s.TxManager(ctx),
			s.WarehouseRepository(ctx),
			s.ReserveRepository(ctx))
	}

	return s.warehouseService
}

func (s *serviceProvider) WarehouseImpl(ctx context.Context) *warehouse.Implementation {
	if s.warehouseImpl == nil {
		s.warehouseImpl = warehouse.NewImplementation(s.WarehouseService(ctx))
	}

	return s.warehouseImpl
}

func setupZeroLog(logConfig *config.LogConfig) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logConfig.TimeFormat}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger = logger.Level(logConfig.LogLevel)
	zerolog.TimeFieldFormat = logConfig.TimeFormat

	return &logger
}
