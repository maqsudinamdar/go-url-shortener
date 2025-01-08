package url

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/url/domain"
	"url-shortener/internal/utils"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload RequestCreateUrl
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if payload.LongURL == "" {
			http.Error(w, "long_url is required", http.StatusBadRequest)
			return
		}

		// Generate a short URL using Base64 encoding
		shortURL := utils.GenerateShortString(payload.LongURL)

		// newUrl
		newUrl := domain.Url{
			ShortURL: shortURL,
			LongURL: payload.LongURL,
		}

		_, err := h.service.Create(newUrl)
		if err != nil {
			HandleError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(newUrl); err != nil {
			HandleError(w, err)
			return
		}

	}
}

func (h *Handler) GetUrlByShortUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := chi.URLParam(r, "shortUrl")
		u, err := h.service.GetUrlByShortUrl(shortUrl)
		if err != nil {
			HandleError(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(u); err != nil {
			HandleError(w, err)
			return
		}
	}
}

