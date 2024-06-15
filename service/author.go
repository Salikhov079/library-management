package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/library-management/genprotos"
	s "github.com/Salikhov079/library-management/storage"
)

type AuthorService struct {
	stg s.StorageRoot
	pb.UnimplementedAuthorServiceServer
}

func NewAuthorService(stg s.StorageRoot) *AuthorService {
	return &AuthorService{stg: stg}
}

func (c *AuthorService) CreateAuthor(ctx context.Context, author *pb.AuthorReq) (*pb.Void, error) {
	res, err := c.stg.Author().Create(author)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *AuthorService) GetAllAuthors(ctx context.Context, author *pb.AuthorReq) (*pb.AllAuthors, error) {
	res, err := c.stg.Author().GetAll(author)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *AuthorService) GetByIdAuthor(ctx context.Context, id *pb.ById) (*pb.AuthorRes, error) {
	res, err := c.stg.Author().GetById(id)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *AuthorService) UpdateAuthor(ctx context.Context, author *pb.AuthorRes) (*pb.Void, error) {
	res, err := c.stg.Author().Update(author)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *AuthorService) DeleteAuthor(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	res, err := c.stg.Author().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return res, err
}
