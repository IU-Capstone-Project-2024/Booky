package booky

import (
	pb "booky-back/api/booky"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	err := s.Storage.DeleteNote(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "DeleteNote: could not delete note: %v", err)
	}

	return &pb.DeleteNoteResponse{}, nil
}
