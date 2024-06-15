package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/library-management/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateBorrower(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	borrower := &pb.BorrowerReq{
		UserId:     "64a04502-985a-44b5-89ec-4bd928c0b402",
		BookId:     "fad3bc35-0235-4b05-90e9-bd11665d4520",
		BorrowDate: "2024-06-06",
		ReturnDate: "2024-06-16",
	}

	result, err := stg.Borrower().Create(borrower)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdBorrower(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "e0b49604-445e-4de2-934d-2161b4f3cb35"

	borrower, err := stg.Borrower().GetById(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, borrower)
}

func TestGetAllBorrwers(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	borrowers, err := stg.Borrower().GetAll(&pb.BorrowerReq{})
	assert.NoError(t, err)
	assert.NotNil(t, borrowers)
}

func TestUpdateBorrower(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	borrower := &pb.BorrowerRes{
		Id:         "e0b49604-445e-4de2-934d-2161b4f3cb35",
		User:       &pb.UserRes{Id: "64a04502-985a-44b5-89ec-4bd928c0b402"},
		Book:       &pb.BookRes{Id: "fad3bc35-0235-4b05-90e9-bd11665d4520"},
		BorrowDate: "2024-06-07",
		ReturnDate: "2024-06-17",
	}
	result, err := stg.Borrower().Update(borrower)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteBorrower(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Borrower().Delete(&Id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
