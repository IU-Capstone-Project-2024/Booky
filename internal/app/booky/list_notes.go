package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/pkg/models"
	"booky-back/internal/pkg/validator"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListNotes(ctx context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	courseID := req.GetCourseId()
	v := validator.New()
	if details, err := v.ValidateID(courseID); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid courseID")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	notes, err := s.Storage.ListNotes(courseID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ListNotes: could not list notes: %v", err)
	}

	grpcNotes := make([]*pb.Note, 0, len(notes))
	for _, n := range notes {
		grpcNote, err := models.BindNoteToGRPC(n)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "ListNotes: could not bind note to grpc: %v", err)
		}
		grpcNotes = append(grpcNotes, grpcNote)
	}

	return &pb.ListNotesResponse{
		Notes: grpcNotes,
	}, nil
}
