package main

import (
	"OMPFinex-CodeChallenge/config"
	"OMPFinex-CodeChallenge/internal/repository"
	"OMPFinex-CodeChallenge/pkg/log"
	rpc "OMPFinex-CodeChallenge/pkg/rpc/proto"
	"OMPFinex-CodeChallenge/services/merger/interactor/merger"
	"OMPFinex-CodeChallenge/services/merger/middleware"
	"OMPFinex-CodeChallenge/services/merger/presenter/grpc_server"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//setting log
	logger := log.NewLogger(log.ErrorLevel, log.NewStdoutCore())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// setting config
	configuration, err := config.NewConfigFromEnv(ctx)
	if err != nil {
		logger.Fatal("can't read config")
	}
	logger.Info("")
	db, err := pgxpool.New(ctx, configuration.Database.Dsn)
	if err != nil {
		logger.Fatal("can't connect database")

	}
	err = db.Ping(ctx)
	if err != nil {
		logger.Fatal("can't connect database")
	}
	logger.Info("it connected to database successfully")

	imageRepo := repository.NewImageRepo(db)
	chunkRepo := repository.NewChunkRepo(db)

	mergerService := merger.New(imageRepo, chunkRepo, logger)

	// Starting Grpc server
	handler := grpc_server.New(mergerService)
	// Interceptors
	grpcOptions := middleware.GRPCInterceptor()

	// Starting gRPC server
	grpcAddress := configuration.GrpcServer.Address
	go runGRPCServer(grpcAddress, grpcOptions, handler, logger)
	logger.Info(fmt.Sprintf("listening on %s/tcp (gRPC)", grpcAddress))
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	<-stopSignal

	logger.Info("stopped Swap server")
}

func runGRPCServer(
	grpcAddress string,
	option grpc.ServerOption,
	server rpc.MergerServer,
	logger log.Logger,
) {
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen on %s: %v", grpcAddress, err))
	}

	grpcServer := grpc.NewServer(option)
	rpc.RegisterMergerServer(grpcServer, server)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
}
