package middleware

import "net/http"

const (
	contentType 	=	"Content-Type"
	applicationJson	= 	"application/json"
)

func CommonMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJson)
		h(w, r)
	}
}