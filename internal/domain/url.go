package domain

// Url domain model
type Url struct {
	ID 			int 	`json:"id" DB:"id"`
	LongURL 	string	`json:"long_url" DB:"long_url"`
	ShortURL	string	`json:"short_url" DB:"short_url"`
}

// Service to manage store
type Service struct {
	s Store
}

func (svc Service) NewService (store Store) *Service {
	return &Service{s: store}
}

func (svc Service) Create(url Url) (int, error) {
	return svc.s.create(url)
}

func (svc Service) GetUrlByShortUrl(shortUrl string) (*Url, error) {
	return svc.s.getUrlByShortUrl(shortUrl)
}
