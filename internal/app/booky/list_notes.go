package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) ListNotes(context.Context, *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	return &pb.ListNotesResponse{}, nil
}
