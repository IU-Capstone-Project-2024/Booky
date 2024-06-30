package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"
)

func (s *Server) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	file, err := s.Storage.GetFile(req.GetId())
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
