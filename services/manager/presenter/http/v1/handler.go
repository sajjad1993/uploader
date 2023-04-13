package v1

import (
	"OMPFinex-CodeChallenge/services/manager/entity"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func (h MangerHandler) registerImage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var image entity.Image
	err := json.NewDecoder(request.Body).Decode(&image)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	err = image.IsValid()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	err = h.downloader.RegisterImage(request.Context(), image)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	h.collector.StartCollecting(request.Context(), image)
	writer.WriteHeader(http.StatusCreated)
}

func (h MangerHandler) createChunk(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var chunk entity.Chunk
	err := json.NewDecoder(request.Body).Decode(&chunk)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	vars := mux.Vars(request)

	sha := vars["sha"]
	chunk.Sha256 = sha
	err = chunk.IsValid()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	err = h.collector.SaveChunk(request.Context(), chunk)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusCreated)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h MangerHandler) getImage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	sha := vars["sha"]
	reader, err := h.uploader.GetImage(request.Context(), sha)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	buf := make([]byte, 1024)
	fmt.Println(" download output")
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			h.logger.Error(err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)

		}
		if n == 0 {
			break
		}
		fmt.Printf("%s", buf)
		// write a chunk
		if _, err := writer.Write(buf[:n]); err != nil {

			h.logger.Error(err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}

}
