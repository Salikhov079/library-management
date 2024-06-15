package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/library-management/genprotos"
	s "github.com/Salikhov079/library-management/storage"
)

type GenreService struct {
	stg s.StorageRoot
	pb.UnimplementedGenreServiceServer
}

func NewGenreService(stg s.StorageRoot) *GenreService {
	return &GenreService{stg: stg}
}

func (c *GenreService) CreateGenre(ctx context.Context, genre *pb.Genre) (*pb.Void, error) {
	res, err := c.stg.Genre().Create(genre)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *GenreService) GetAllGenres(ctx context.Context, genre *pb.Genre) (*pb.AllGenres, error) {
	res, err := c.stg.Genre().GetAll(genre)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *GenreService) GetByIdGenre(ctx context.Context, id *pb.ById) (*pb.Genre, error) {
	prod, err := c.stg.Genre().GetById(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *GenreService) UpdateGenre(ctx context.Context, genre *pb.Genre) (*pb.Void, error) {
	res, err := c.stg.Genre().Update(genre)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (c *GenreService) DeleteGenre(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	res, err := c.stg.Genre().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return res, err
}
