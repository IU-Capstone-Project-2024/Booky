package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/storage"
	inmemory "booky-back/internal/storage/in-memory"
)

type Server struct {
	pb.UnimplementedBookyServiceServer
	Storage storage.DataModel
}

func NewServer() *Server {
	return &Server{
		Storage: inmemory.NewInMemoryStorage(),
	}
}
