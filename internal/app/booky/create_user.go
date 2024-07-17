package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/pkg/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := models.BindUser(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateUser: could not bind user: %v", err)
	}

	if !user.Validate() {
		return nil, status.Error(codes.InvalidArgument, "CreateUser: user is not valid")
	}

	returnedUser, err := s.Storage.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateUser: could not create user: %v", err)
	}

	grpcUser, err := models.BindUserToGRPC(returnedUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateUser: could not bind user to grpc: %v", err)
	}

	return &pb.CreateUserResponse{
		User: grpcUser,
	}, nil
}
