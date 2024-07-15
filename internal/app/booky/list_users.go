package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.Storage.ListUsers()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ListUsers: could not list users: %v", err)
	}

	grpcUsers := make([]*pb.User, 0, len(users))
	for _, user := range users {
		grpcUser, err := models.BindUserToGRPC(user)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "ListUsers: could not bind user to grpc: %v", err)
		}

		grpcUsers = append(grpcUsers, grpcUser)
	}

	return &pb.ListUsersResponse{
		Users: grpcUsers,
	}, nil
}
