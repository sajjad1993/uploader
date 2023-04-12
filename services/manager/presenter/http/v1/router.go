package v1

import (
	"OMPFinex-CodeChallenge/pkg/log"
	"OMPFinex-CodeChallenge/services/manager/interactor/collector"
	"OMPFinex-CodeChallenge/services/manager/interactor/downloader"
	"github.com/gorilla/mux"
	"net/http"
)

type MangerHandler struct {
	downloader donwloder.UseCase
	collector  collector.UseCase
	logger     log.Logger
}

func NewMangerHandler(downloader donwloder.UseCase, collector collector.UseCase, logger log.Logger) http.Handler {
	handler := MangerHandler{
		downloader: downloader,
		collector:  collector,
		logger:     logger,
	}
	router := mux.NewRouter()

	router.HandleFunc("/image", handler.registerImage).Methods(http.MethodPost)
	router.HandleFunc("/image/{sha}/chunks", handler.createChunk).Methods(http.MethodPost)
	router.HandleFunc("/image/{sha}", handler.getImage).Methods(http.MethodGet)
	return router
}
