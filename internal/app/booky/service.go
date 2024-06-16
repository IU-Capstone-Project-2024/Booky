package booky

import (
	pb "booky-back/api/booky"
)

type Server struct {
	pb.UnimplementedBookyServiceServer
}

func NewServer() *Server {
	return &Server{}
}
