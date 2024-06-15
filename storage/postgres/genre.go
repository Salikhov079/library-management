package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/google/uuid"
)

type GenreStorage struct {
	db *sql.DB
}

func NewGenreStorage(db *sql.DB) *GenreStorage {
	return &GenreStorage{db: db}
}

func (p *GenreStorage) Create(genre *pb.Genre) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO genres 
			(id, name)
		VALUES 
			($1, $2)
	`
	_, err := p.db.Exec(query, id, genre.Name)
	return nil, err
}

func (p *GenreStorage) GetById(id *pb.ById) (*pb.Genre, error) {
	query := `
		SELECT name
		FROM genres
		WHERE id = $1 and deleted_at = 0
	`

	row := p.db.QueryRow(query, id.Id)
	var genre pb.Genre

	err := row.Scan(
		&genre.Name)
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (p *GenreStorage) GetAll(genre *pb.Genre) (*pb.AllGenres, error) {
	genres := &pb.AllGenres{}
	var arr []interface{}
	count := 1

	query := `
		SELECT name
		FROM genres
		WHERE deleted_at = 0 
	`

	if len(genre.Name) > 0 {
		query += fmt.Sprintf(" and name=$%d ", count)
		count++
		arr = append(arr, genre.Name)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var genre pb.Genre
		err := row.Scan(
			&genre.Name)
		if err != nil {
			return nil, err
		}
		genres.Genres = append(genres.Genres, &genre)
	}

	return genres, nil
}

func (p *GenreStorage) Update(genre *pb.Genre) (*pb.Void, error) {
	query := `
		UPDATE genres
		SET name = $2, updated_at = now()
		WHERE id = $1 and deleted_at = 0 
	`
	_, err := p.db.Exec(query, genre.Id, genre.Name)
	return nil, err
}

func (p *GenreStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE genres 
		SET deleted_at = $2
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}
