package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) CreateNote(context.Context, *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	return &pb.CreateNoteResponse{}, nil
}
