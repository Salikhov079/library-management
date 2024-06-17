package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	book := &pb.BookReq{
		Title:    "Bir kun",
		AuthorId: "de041c42-9373-4ef0-aaf7-f107f460abfd",
		GenreId:  "7a55f1a2-ac9d-4a5c-bab3-9a823ed328ba",
		Summary:  "Hello World",
	}

	result, err := stg.Book().Create(book)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdBook(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "fad3bc35-0235-4b05-90e9-bd11665d4520"

	book, err := stg.Book().GetById(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestGetAllBook(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	books, err := stg.Book().GetAll(&pb.FilterBook{})
	assert.NoError(t, err)
	assert.NotNil(t, books)
}

func TestUpdateBook(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	book := &pb.BookRes{
		Id: "fad3bc35-0235-4b05-90e9-bd11665d4520",
		Title:    "O'sha kun",
		Author: &pb.AuthorRes{Id: "de041c42-9373-4ef0-aaf7-f107f460abfd"},
		Genre:  &pb.Genre{Id: "7a55f1a2-ac9d-4a5c-bab3-9a823ed328ba"},
		Summary:  "World",
	}
	result, err := stg.Book().Update(book)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteBook(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Book().Delete(&Id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
