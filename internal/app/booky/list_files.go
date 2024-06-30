package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	files, err := s.Storage.ListFiles(req.GetCourseId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ListFiles: could not list files: %v", err)
	}

	grpcFiles := make([]*pb.File, 0, len(files))
	for _, file := range files {
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
