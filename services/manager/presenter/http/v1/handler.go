package v1

import (
	"OMPFinex-CodeChallenge/services/manager/entity"
	"encoding/json"
	"net/http"
)

func (h MangerHandler) registerImage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var image entity.Image
	err := json.NewDecoder(request.Body).Decode(&image)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = h.downloader.RegisterImage(request.Context(), image)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h MangerHandler) createChunk(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var chunk entity.Chunk
	err := json.NewDecoder(request.Body).Decode(&chunk)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = h.collector.SaveChunk(request.Context(), chunk)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h MangerHandler) getImage(writer http.ResponseWriter, request *http.Request) {

}
