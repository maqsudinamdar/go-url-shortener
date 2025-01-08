package url

import (
	"url-shortener/cmd/middleware"
	"url-shortener/internal/url/domain"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	Create(url domain.Url) (int, error)
	GetUrlByShortUrl(shortUrl string) (*domain.Url, error)
}

type Router struct {
	r 		*chi.Mux
	service Service
}

func NewRouter(r *chi.Mux, db *sqlx.DB) *Router {
	return &Router{
		r: 			r,
		service: 	domain.NewService(domain.NewStore(db))}
}

func (h *Router) Routes() {
	handler := NewHandler(h.service)
	h.r.Post("/api/v1/url/", middleware.CommonMiddleware(handler.CreateUrl()))
	h.r.Get("/api/v1/url/{shortUrl:[a-zA-Z0-9]{6}}", middleware.CommonMiddleware(handler.GetUrlByShortUrl()))
}