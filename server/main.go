package main

import (
	"log"
	"net"

	pb "github.com/Salikhov079/library-management/genprotos"
	"github.com/Salikhov079/library-management/service"
	"github.com/Salikhov079/library-management/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	liss, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterAuthorServiceServer(s, service.NewAuthorService(db))
	pb.RegisterBookServiceServer(s, service.NewBookService(db))
	pb.RegisterGenreServiceServer(s, service.NewGenreService(db))
	

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
