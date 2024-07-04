package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"booky-back/internal/storage"
	"booky-back/internal/validator"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	id := req.GetId()
	v := validator.New()
	if details, err := v.ValidateID(id); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid id")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	note, err := s.Storage.GetNote(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "GetNote: note not found")
		}
		return nil, status.Errorf(codes.Internal, "GetNote: could not get note: %v", err)
	}

	grpcNote, err := models.BindNoteToGRPC(note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetNote: could not bind note to grpc: %v", err)
	}

	return &pb.GetNoteResponse{
		Note: grpcNote,
	}, nil
}
