package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) UpdateNote(context.Context, *pb.UpdateNoteRequest) (*pb.UpdateNoteResponse, error) {
	return &pb.UpdateNoteResponse{}, nil
}
