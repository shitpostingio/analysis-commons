package handler

import (
	"mime/multipart"
	"net/http"
	"strings"
)

func Download(maxSize int64, r *http.Request) (multipart.File, string, int, error) {

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		return nil, "", http.StatusBadRequest, err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, "", http.StatusBadRequest, err
	}

	extensions := strings.Split(header.Filename, ".")
	return file, extensions[len(extensions)-1], 0, nil

}
