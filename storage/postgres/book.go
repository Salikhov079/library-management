package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/google/uuid"
)

type BookStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (p *BookStorage) Create(book *pb.BookReq) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO books 
			(id, title, author_id, genre_id, summary)
		VALUES 
			($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, book.Title, book.AuthorId, book.GenreId, book.Summary)
	return nil, err
}

func (p *BookStorage) GetById(id *pb.ById) (*pb.BookRes, error) {
	query := `
			SELECT 
				b.title, a.name, a.biography, g.name, b.summary
			FROM 
				books b JOIN authors a
			ON 
				b.author_id = a.id
			JOIN genres g ON
				b.genre_id = g.id
			WHERE 
				b.id = $1 and b.deleted_at = 0 
		`

	row := p.db.QueryRow(query, id.Id)
	var book pb.BookRes
	var author pb.AuthorRes
	var genre pb.Genre

	err := row.Scan(
		&book.Title,
		&author.Name,
		&author.Biography,
		&genre.Name,
		&book.Summary)
	if err != nil {
		return nil, err
	}
	book.Author = &author
	book.Genre = &genre

	return &book, nil
}

func (p *BookStorage) GetAll(book *pb.BookReq) (*pb.AllBooks, error) {
	books := &pb.AllBooks{}
	var arr []interface{}
	count := 1

	query := `
			SELECT 
				b.title, a.name, a.biography, g.name, b.summary
			FROM 
				books b JOIN authors a
			ON 
				b.author_id = a.id
			JOIN genres g ON
				b.genre_id = g.id
			WHERE 
				b.deleted_at = 0
		`

	if len(book.Title) > 0 {
		query += fmt.Sprintf(" and title=$%d ", count)
		count++
		arr = append(arr, book.Title)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var book pb.BookRes
		var author pb.AuthorRes
		var genre pb.Genre

		err := row.Scan(
			&book.Title,
			&author.Name,
			&author.Biography,
			&genre.Name,
			&book.Summary)
		if err != nil {
			return nil, err
		}

		book.Author = &author
		book.Genre = &genre
		books.Books = append(books.Books, &book)
	}

	return books, nil
}

func (p *BookStorage) Update(book *pb.BookRes) (*pb.Void, error) {
	query := `
		UPDATE books
		SET title = $2, author_id = $3, genre_id = $4, summary = $5, updated_at = now()
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, book.Id ,book.Title, book.Author.Id, book.Genre.Id, book.Summary)
	return nil, err
}

func (p *BookStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE books 
		SET deleted_at = $2
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}
