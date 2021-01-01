package downloader

import (
	"io"
	"net/http"
)

const (
	callerAPIKeyHeaderName string = "X-caller-bot-apikey"
	downloadURLHeaderName  string = "X-download-file-url"
)

// Downloader represents a generic handler for file downloads.
type Downloader interface {
	Download(id, mediaType string, maxSize int64, r *http.Request) (filename string, reader io.ReadCloser, err error)
}
