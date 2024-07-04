package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"booky-back/internal/validator"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	id := req.GetId()
	v := validator.New()
	if details, err := v.ValidateID(id); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid id")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	file, err := s.Storage.GetFile(id)
	if err != nil {
		return nil, err
	}

	grpcFile, err := models.BindFileToGRPC(file)
	if err != nil {
		return nil, err
	}

	return &pb.GetFileResponse{
		File: grpcFile,
	}, nil
}
