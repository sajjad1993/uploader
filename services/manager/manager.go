package main

import (
	"OMPFinex-CodeChallenge/config"
	"OMPFinex-CodeChallenge/internal/repository"
	"OMPFinex-CodeChallenge/pkg/log"
	"OMPFinex-CodeChallenge/services/manager/interactor/collector"
	"OMPFinex-CodeChallenge/services/manager/interactor/downloader"
	v1 "OMPFinex-CodeChallenge/services/manager/presenter/http/v1"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"

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

	collectorService, err := collector.New(imageRepo, chunkRepo, configuration.GrpcServer.Address, logger)
	if err != nil {
		logger.Fatal("can't connect grpc")
	}
	go collectorService.CallMerger()
	downloaderService := donwloder.New(imageRepo, collectorService.GetChannel(), logger)

	// Starting HTTP server
	handler := v1.NewMangerHandler(downloaderService, collectorService, logger)
	go runHTTPServer(handler,
		configuration.HTTPServer.Address,
		configuration.HTTPServer.Port,
		configuration.HTTPServer.ReadTimeout,
		configuration.HTTPServer.WriteTimeout, logger,
	)
	logger.Info(fmt.Sprintf("listening on %s (http)", configuration.HTTPServer.Address))

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	<-stopSignal

	logger.Info("stopped Swap server")
}

func runHTTPServer(
	handler http.Handler,
	address string,
	port uint,
	writeTimeout time.Duration,
	readTimeout time.Duration,
	logger log.Logger,
) {
	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%s:%d", address, port),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	logger.Fatal(srv.ListenAndServe().Error())
}
