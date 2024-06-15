package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Salikhov079/library-management/config"
	u "github.com/Salikhov079/library-management/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db        *sql.DB
	Books     u.Book
	Authors   u.Author
	Genres    u.Genre
	Borrowers u.Borrower
}

func NewPostgresStorage() (u.StorageRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, err

}

func (s *Storage) Book() u.Book {
	if s.Books == nil {
		s.Books = &BookStorage{db: s.Db}
	}
	return s.Books
}

func (s *Storage) Author() u.Author {
	if s.Authors == nil {
		s.Authors = &AuthorStorage{db: s.Db}
	}
	return s.Authors
}

func (s *Storage) Genre() u.Genre {
	if s.Genres == nil {
		s.Genres = &GenreStorage{db: s.Db}
	}
	return s.Genres
}

func (s *Storage) Borrower() u.Borrower {
	if s.Borrowers == nil {
		s.Borrowers = &BorrowerStorage{db: s.Db}
	}
	return s.Borrowers
}
