package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) DeleteNote(context.Context, *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	return &pb.DeleteNoteResponse{}, nil
}
