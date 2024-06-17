package postgres

import (
	pb "github.com/Salikhov079/library-management/genprotos"
)

type StorageRoot interface {
	Book() Book
	Author() Author
	Genre() Genre
	Borrower() Borrower
}

type Book interface {
	Create(book *pb.BookReq) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.BookRes, error)
	GetAll(filter *pb.FilterBook) (*pb.AllBooks, error)
	Update(book *pb.BookRes) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}

type Author interface {
	Create(author *pb.AuthorReq) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.AuthorRes, error)
	GetAll(filter *pb.FilterAuthor) (*pb.AllAuthors, error)
	Update(author *pb.AuthorRes) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}

type Genre interface {
	Create(author *pb.Genre) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.Genre, error)
	GetAll(filter *pb.FilterGenre) (*pb.AllGenres, error)
	Update(author *pb.Genre) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}

type Borrower interface {
	Create(author *pb.BorrowerReq) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.BorrowerRes, error)
	GetAll(filter *pb.FilterBorrower) (*pb.AllBorrowers, error)
	Update(author *pb.BorrowerRes) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	GetAllId(*pb.Void) (*pb.AllBorrowers, error)
}
