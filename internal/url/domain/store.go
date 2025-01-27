package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//Store interface for Url persistence layer
type Store interface {
	create(url Url) (int, error)
	getUrlByShortUrl(shortUrl string) (*Url, error)
}

// UrlStore for persistance
type UrlStore struct {
	db *sqlx.DB
}

//NewStore creates new UrlStore for Urls
func NewStore(db *sqlx.DB) *UrlStore {
	return &UrlStore{db: db}
}


func (s *UrlStore) create(url Url) (int, error) {
	result, err := s.db.Exec("INSERT INTO url (long_url, short_url) VALUES (?, ?)", url.LongURL, url.ShortURL)
	if err != nil {
		return 0, ErrDbQuery{Err: errors.Wrap(err, "")}
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, ErrDbNotSupported{Err:  errors.Wrap(err, "UrlStore.create() error")}
	}
	return int(lastId), nil
}

func (s *UrlStore) getUrlByShortUrl(shortUrl string) (*Url, error) {
	var url Url
	err := s.db.Get(&url, "SELECT * FROM url WHERE short_url=?", shortUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEntityNotExist{Err: errors.Wrap(err, "UrlStore.getUrlByShortUrl() ErrNoRows error")}
		}
		return nil, ErrDbQuery{Err: errors.Wrap(err, "")}
	}
	return &url, nil
}