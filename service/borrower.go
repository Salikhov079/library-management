package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/library-management/genprotos"
	s "github.com/Salikhov079/library-management/storage"
)

type BorrowerService struct {
	stg s.StorageRoot
	pb.UnimplementedBorrowerServiceServer
}

func NewBorrowerService(stg s.StorageRoot) *BorrowerService {
	return &BorrowerService{stg: stg}
}

func (c *BorrowerService) CreateBorrower(ctx context.Context, borrower *pb.BorrowerReq) (*pb.Void, error) {
	res, err := c.stg.Borrower().Create(borrower)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *BorrowerService) GetAllBorrowers(ctx context.Context, borrower *pb.BorrowerReq) (*pb.AllBorrowers, error) {
	res, err := c.stg.Borrower().GetAll(borrower)
	if err != nil {
		log.Print(err)
	}
	


	return res, err
}

func (c *BorrowerService) GetByIdBorrower(ctx context.Context, id *pb.ById) (*pb.BorrowerRes, error) {
	res, err := c.stg.Borrower().GetById(id)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *BorrowerService) UpdateBorrower(ctx context.Context, Borrower *pb.BorrowerRes) (*pb.Void, error) {
	res, err := c.stg.Borrower().Update(Borrower)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *BorrowerService) DeleteBorrower(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	res, err := c.stg.Borrower().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return res, err
}
