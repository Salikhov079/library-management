package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/library-management/genprotos"
	s "github.com/Salikhov079/library-management/storage"
)

type BookService struct {
	stg s.StorageRoot
	pb.UnimplementedBookServiceServer
}

func NewBookService(stg s.StorageRoot) *BookService {
	return &BookService{stg: stg}
}

func (c *BookService) CreateBook(ctx context.Context, book *pb.BookReq) (*pb.Void, error) {
	res, err := c.stg.Book().Create(book)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *BookService) GetAllBooks(ctx context.Context, book *pb.BookReq) (*pb.AllBooks, error) {
	res, err := c.stg.Book().GetAll(book)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *BookService) GetByIdBook(ctx context.Context, id *pb.ById) (*pb.BookRes, error) {
	prod, err := c.stg.Book().GetById(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *BookService) UpdateBook(ctx context.Context, book *pb.BookRes) (*pb.Void, error) {
	res, err := c.stg.Book().Update(book)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *BookService) DeleteBook(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	res, err := c.stg.Book().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return res, err
}
