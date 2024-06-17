package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateGenre(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	genre := &pb.Genre{
		Name: "Jangari",
	}

	result, err := stg.Genre().Create(genre)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdGenre(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "7a55f1a2-ac9d-4a5c-bab3-9a823ed328ba"

	genre, err := stg.Genre().GetById(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, genre)
}

func TestGetAllGenre(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	genres, err := stg.Genre().GetAll(&pb.FilterGenre{})
	assert.NoError(t, err)
	assert.NotNil(t, genres)
}

func TestUpdateGenre(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	genre := &pb.Genre{
		Id:   "7a55f1a2-ac9d-4a5c-bab3-9a823ed328ba",
		Name: "Komedik",
	}
	result, err := stg.Genre().Update(genre)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteGenre(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Genre().Delete(&Id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
