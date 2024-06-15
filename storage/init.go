package postgres

import (
	pb "github.com/Salikhov079/library-management/genprotos"
)

type StorageRoot interface {
	Book() Book
	Author() Author
	Genre() Genre
}

type Book interface {
	Create(book *pb.BookReq) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.BookRes, error)
	GetAll(_ *pb.BookReq) (*pb.AllBooks, error)
	Update(book *pb.BookRes) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}

type Author interface {
	Create(author *pb.AuthorReq) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.AuthorRes, error)
	GetAll(_ *pb.AuthorReq) (*pb.AllAuthors, error)
	Update(author *pb.AuthorRes) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}

type Genre interface {
	Create(author *pb.Genre) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.Genre, error)
	GetAll(_ *pb.Genre) (*pb.AllGenres, error)
	Update(author *pb.Genre) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}