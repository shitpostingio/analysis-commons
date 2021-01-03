package downloader

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// TelegramDownloader implements the download for telegram files.
type DirectDownloader struct{
	// MaxDownloadSize represents the maximum allowed download size for medias
	MaxDownloadSize int64
}

// Download downloads telegram files.
func (d *DirectDownloader) Download(id, mediaType string, r *http.Request) (filename string, reader io.ReadCloser, err error) {

	file, extension, _, err := download(d.GetMaxDownloadSize(), r)
	if extension == "" {

		if mediaType == "image" {
			extension = "jpg"
		} else {
			extension = "mp4"
		}

	}

	filename = fmt.Sprintf("%s.%s", id, extension)
	reader = file
	return

}

func download(maxSize int64, r *http.Request) (multipart.File, string, int, error) {

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

func (d *DirectDownloader) GetMaxDownloadSize() int64 {
	return d.MaxDownloadSize
}
