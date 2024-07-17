package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/pkg/models"
	"booky-back/internal/pkg/validator"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFile(ctx context.Context, req *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {
	file, err := models.BindFile(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "CreateFile: could not bind file: %v", err)
	}

	v := validator.New()
	if details, err := v.ValidateFile(file); err != nil {
		return nil, status.Error(codes.InvalidArgument, "CreateFile: validation error: invalid file data")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "CreateFile: validation error: %v", details)
	}

	user, err := s.Storage.GetUser(file.Publisher.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "CreateFile: could not get user: %v", err)
	}

	file.Publisher = *user

	createdFile, err := s.Storage.CreateFile(file)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateFile: could not create file: %v", err)
	}

	grpcFile, err := models.BindFileToGRPC(createdFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateFile: could not bind file to grpc: %v", err)
	}

	return &pb.CreateFileResponse{
		File: grpcFile,
	}, nil
}
