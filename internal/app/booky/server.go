package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/config"
	"booky-back/internal/gpt"
	"booky-back/internal/storage"
	inmemory "booky-back/internal/storage/in-memory"
)

type Server struct {
	Config  *config.Config
	Storage storage.DataModel
	GPT     *gpt.GPT

	pb.UnimplementedBookyServiceServer
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config:  config,
		Storage: inmemory.NewInMemoryStorage(),
	}
}
