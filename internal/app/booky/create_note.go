package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	note, err := models.BindNote(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "CreateNote: could not bind note: %v", err)
	}

	user, err := s.Storage.GetUser(note.Publisher.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "CreateNote: could not get user: %v", err)
	}

	note.Publisher = *user

	createdNote, err := s.Storage.CreateNote(note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateNote: could not create note: %v", err)
	}

	grpcNote, err := models.BindNoteToGRPC(createdNote)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateNote: could not bind note to grpc: %v", err)
	}

	return &pb.CreateNoteResponse{
		Note: grpcNote,
	}, nil
}
