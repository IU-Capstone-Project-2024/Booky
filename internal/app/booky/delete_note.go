package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/storage"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	err := s.Storage.DeleteNote(req.GetId())
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "DeleteNote: note not found")
		}
		return nil, status.Errorf(codes.Internal, "DeleteNote: could not delete note: %v", err)
	}

	return &pb.DeleteNoteResponse{}, nil
}
