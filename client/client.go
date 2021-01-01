package client

import (
	"bytes"
	"encoding/json"
	"github.com/shitpostingio/analysis-commons/structs"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// PerformRequest performs a request to the analysis service.
func PerformRequest(file io.Reader, fileName, endpoint, authorizationHeaderName, authorizationHeaderValue string) (data structs.Analysis, err error) {

	//
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Send files via multipart/form-data
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	err = writer.Close()
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return
	}

	//
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add(authorizationHeaderName, authorizationHeaderValue)
	client := http.Client{Timeout: time.Second * 30}
	response, err := client.Do(request)
	if err != nil {
		return
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println("PerformRequest: unable to close response body", err)
		}
	}()

	//
	bodyResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	log.Debugln("Request result: ", string(bodyResult))

	//
	var ar structs.Analysis
	err = json.Unmarshal(bodyResult, &ar)
	if err != nil {
		log.Println("PerformRequest: error while unmarshaling ", err)
		return
	}

	return ar, err

}
