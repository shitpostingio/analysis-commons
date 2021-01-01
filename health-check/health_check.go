package health_check

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ConfirmServiceHealth sends a simple string to confirm service health.
func ConfirmServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := fmt.Fprintln(w, "😂👌")
	if err != nil {
		log.Println("Unable to respond to health check: ", err)
	}
}
