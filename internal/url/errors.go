package url

import (
	"encoding/json"
	"net/http"
	"strings"
	"url-shortener/internal/url/domain"

	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JsonError(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(ErrorResponse{error}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleError(w http.ResponseWriter, err error) {
	const logFormat = "fatal: %+v\n"
	if strings.Contains(err.Error(), "connection refused") {
		log.Warnf(logFormat, err)
		JsonError(w, "DB_CONNECTION_FAIL", http.StatusServiceUnavailable)
		return
	}
	if err.Error() == http.StatusText(400) {
		log.Warnf(logFormat, err)
		JsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch err.(type) {
	case domain.ErrDbQuery:
		log.Warnf(logFormat, err.(domain.ErrDbQuery).Err)
		JsonError(w, err.Error(), http.StatusConflict)
	case domain.ErrDbNotSupported:
		log.Warnf(logFormat, err.(domain.ErrDbNotSupported).Err)
		JsonError(w, err.Error(), http.StatusConflict)
	case domain.ErrEntityNotExist:
		log.Warnf(logFormat, err.(domain.ErrEntityNotExist).Err)
		JsonError(w, err.Error(), http.StatusNotFound)
	default:
		log.Warnf(logFormat, err)
		JsonError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}