package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"booky-back/internal/validator"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	courseID := req.GetCourseId()
	v := validator.New()
	if details, err := v.ValidateID(courseID); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid courseID")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	files, err := s.Storage.ListFiles(courseID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ListFiles: could not list files: %v", err)
	}

	grpcFiles := make([]*pb.File, 0, len(files))
	for _, file := range files {
		user, err := s.Storage.GetUser(file.Publisher.ID)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "CreateFile: could not get user: %v", err)
		}

		file.Publisher = *user

		grpcFile, err := models.BindFileToGRPC(file)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "ListFiles: could not bind file to grpc: %v", err)
		}

		grpcFiles = append(grpcFiles, grpcFile)
	}

	return &pb.ListFilesResponse{
		Files: grpcFiles,
	}, nil
}
