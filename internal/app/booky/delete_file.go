package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	err := s.Storage.DeleteFile(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteFileResponse{}, nil
}
