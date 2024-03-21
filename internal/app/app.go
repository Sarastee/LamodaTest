package app

import (
	"context"
	"log"
	"net"

	"github.com/sarastee/LamodaTest/internal/config"
	desc "github.com/sarastee/LamodaTest/pkg/warehouse_v1"
	"github.com/sarastee/platform_common/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App is struct, containing pointer for serviceProvider, pointer for grpc.Server, pointer for *http.Server and configPath
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	configPath      string
}

// NewApp is a method, returning pointer for App struct
func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run is a method, starting gRPC server and defer CloseAll function
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	initDepFunctions := []func(ctx2 context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range initDepFunctions {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterWarehouseV1Server(a.grpcServer, a.serviceProvider.WarehouseImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC started on %s", a.serviceProvider.GRPCConfig().Address())
	listener, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
