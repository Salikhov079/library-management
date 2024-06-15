package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/google/uuid"
)

type AuthorStorage struct {
	db *sql.DB
}

func NewAuthorStorage(db *sql.DB) *AuthorStorage {
	return &AuthorStorage{db: db}
}

func (p *AuthorStorage) Create(author *pb.AuthorReq) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO authors 
			(id, name, biography)
		VALUES 
			($1, $2, $3)
	`
	_, err := p.db.Exec(query, id, author.Name, author.Biography)
	return nil, err
}

func (p *AuthorStorage) GetById(id *pb.ById) (*pb.AuthorRes, error) {
	query := `
		SELECT name, biography
		FROM authors
		WHERE id = $1 and deleted_at = 0
	`

	row := p.db.QueryRow(query, id.Id)
	var author pb.AuthorRes

	err := row.Scan(
		&author.Name,
		&author.Biography)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (p *AuthorStorage) GetAll(author *pb.AuthorReq) (*pb.AllAuthors, error) {
	authors := &pb.AllAuthors{}
	var arr []interface{}
	count := 1

	query := `
		SELECT name, biography
		FROM authors
		WHERE deleted_at = 0 
	`

	if len(author.Name) > 0 {
		query += fmt.Sprintf(" and name=$%d ", count)
		count++
		arr = append(arr, author.Name)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var author pb.AuthorRes
		err := row.Scan(
			&author.Name,
			&author.Biography)
		if err != nil {
			return nil, err
		}
		authors.Authors = append(authors.Authors, &author)
	}

	return authors, nil
}

func (p *AuthorStorage) Update(author *pb.AuthorRes) (*pb.Void, error) {
	query := `
		UPDATE authors
		SET name = $2, biography = $3, updated_at = now()
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, author.Id, author.Name, author.Biography)
	return nil, err
}

func (p *AuthorStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE authors 
		SET deleted_at = $2
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}
