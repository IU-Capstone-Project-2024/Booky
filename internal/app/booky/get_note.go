package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) GetNote(context.Context, *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	return &pb.GetNoteResponse{}, nil
}
