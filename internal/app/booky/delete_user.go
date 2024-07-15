package booky

import (
	pb "booky-back/api/booky"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.Storage.DeleteUser(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "DeleteUser: could not delete user: %v", err)
	}

	return &pb.DeleteUserResponse{}, nil
}
