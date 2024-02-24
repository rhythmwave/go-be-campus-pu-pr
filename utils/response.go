package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	contentType              = "Content-Type"
	contentTypeValue         = "application/json; charset=utf-8"
	xContentTypeOptions      = "X-Content-Type-Options"
	xContentTypeOptionsValue = "nosniff"
)

func JSONResponse(w http.ResponseWriter, statusCode int, r interface{}) {
	w.Header().Set(contentType, contentTypeValue)
	w.Header().Set(xContentTypeOptions, xContentTypeOptionsValue)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Error(err)
	}
}
