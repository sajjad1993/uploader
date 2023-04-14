package main

import (
	"OMPFinex-CodeChallenge/config"
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/internal/repository"
	"OMPFinex-CodeChallenge/pkg/log"
	"OMPFinex-CodeChallenge/services/manager/interactor/collector"
	"OMPFinex-CodeChallenge/services/manager/interactor/downloader"
	"OMPFinex-CodeChallenge/services/manager/interactor/uploader"
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
	logger.Info(fmt.Sprintf("download-interval = %s | retry = %d ", configuration.DownloadImage.RetryInterval, configuration.DownloadImage.Retry))
	if err != nil {
		logger.Fatal("can't read config")
	}

	// create repository instance
	var imageRepo image.Repository
	var chunkRepo chunk.Repository

	if configuration.Database.Disable {
		logger.Info("create none database repository")

		imageRepo, chunkRepo, err = noneDatabaseRepository(configuration, logger)
		if err != nil {
			logger.Fatal("can't connect repository")
		}
	} else {
		logger.Info("create  database repository")

		imageRepo, chunkRepo, err = databaseRepository(logger, ctx, configuration)
		if err != nil {
			logger.Fatal("can't connect repository")
		}
	}

	collectorService, err := collector.New(imageRepo, chunkRepo, configuration.GrpcServer.Address, logger, configuration.GlobalTimeOut)
	if err != nil {
		logger.Fatal("can't connect grpc")
	}
	go collectorService.CallMerger()
	downloaderService := donwloder.New(imageRepo, collectorService.GetChannel(), logger)
	uploaderService := uploader.New(imageRepo, configuration.DownloadImage.Retry, configuration.DownloadImage.RetryInterval, logger)
	// Starting HTTP server
	handler := v1.NewMangerHandler(downloaderService, collectorService, uploaderService, logger)
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

func databaseRepository(logger log.Logger, ctx context.Context, configuration config.Config) (image.Repository, chunk.Repository, error) {
	logger.Info("")
	db, err := pgxpool.New(ctx, configuration.Database.Dsn)
	if err != nil {
		return nil, nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, nil, err
	}
	logger.Info("it connected to database successfully")
	imageRepo := repository.NewImageRepo(db)
	chunkRepo := repository.NewChunkRepo(db)
	return imageRepo, chunkRepo, nil
}

func noneDatabaseRepository(configuration config.Config, logger log.Logger) (image.Repository, chunk.Repository, error) {
	storageConfig := configuration.Storage
	imageRepo, err := repository.NewImageMemory(storageConfig.Images, storageConfig.Chunks)
	if err != nil {
		return nil, nil, err
	}
	chunkRepo, err := repository.NewChunkMemory(storageConfig.Images, storageConfig.Chunks)
	if err != nil {
		return nil, nil, err
	}
	return imageRepo, chunkRepo, nil
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
