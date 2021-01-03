package downloader

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// TelegramDownloader implements the download for telegram files.
type TelegramDownloader struct{
	// MaxDownloadSize represents the maximum allowed download size for medias
	MaxDownloadSize int64
}

// Download downloads telegram files.
func (*TelegramDownloader) Download(id, mediaType string, r *http.Request) (filename string, reader io.ReadCloser, err error) {

	// Filename
	workingFileURL := r.Header.Get(downloadURLHeaderName)
	words := strings.Split(workingFileURL, ".")
	extension := words[len(words)-1]
	if extension == "" {

		if mediaType == "image" {
			extension = "jpg"
		} else {
			extension = "mp4"
		}

	}

	filename = fmt.Sprintf("%s.%s", id, extension)

	// Reader
	callerTgAPIKey := r.Header.Get(callerAPIKeyHeaderName)
	workingFileURL = fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", callerTgAPIKey, workingFileURL)
	resp, err := http.Get(workingFileURL)
	if err != nil {
		return
	}

	reader = resp.Body
	return

}

func (d *TelegramDownloader) GetMaxDownloadSize() int64 {
	return d.MaxDownloadSize
}
