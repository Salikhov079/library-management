package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	author := &pb.AuthorReq{
		Name: "Abdulla Qodiriy",
		Biography: "Salomat",
	}

	result, err := stg.Author().Create(author)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdAuthor(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "de041c42-9373-4ef0-aaf7-f107f460abfd"

	author, err := stg.Author().GetById(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, author)
}

func TestGetAllAuthor(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	authors, err := stg.Author().GetAll(&pb.AuthorReq{})
	assert.NoError(t, err)
	assert.NotNil(t, authors)
}

func TestUpdateAuthor(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	author := &pb.AuthorRes{
		Id: "de041c42-9373-4ef0-aaf7-f107f460abfd",
		Name: "Abdulla Qodiriy",
		Biography: "Exe Salomat",
	}
	result, err := stg.Author().Update(author)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteAuthor(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Author().Delete(&Id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
