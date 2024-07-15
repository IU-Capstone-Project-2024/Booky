package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.Storage.GetUser(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetUser: could not get user: %v", err)
	}

	grpcUser, err := models.BindUserToGRPC(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetUser: could not bind user to grpc: %v", err)
	}

	return &pb.GetUserResponse{
		User: grpcUser,
	}, nil
}
