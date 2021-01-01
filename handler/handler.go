package handler

import (
	"encoding/json"
	"github.com/shitpostingio/analysis-commons/decoder"
	"github.com/shitpostingio/analysis-commons/structs"
	"io"
	"log"
	"net/http"
)

type Handler func(string, io.Reader, decoder.MediaDecoder) *structs.Analysis

// Handle is a generic Handler that downloads a file, performs an analysis function and sends the response back.
func Handle(w http.ResponseWriter, r *http.Request, maxSize int64, decoder decoder.MediaDecoder, h Handler) {

	file, extension, status, err := Download(maxSize, r)
	if err != nil {
		http.Error(w, "error while downloading media", status)
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("Handle: unable to close request data ", err)
		}
	}()

	result := h(extension, file, decoder)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Handle: unable to send response: ", err)
	}

}
