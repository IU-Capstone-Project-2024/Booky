package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	course, err := s.Storage.GetCourse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetCourse: failed to get course: %v", err)
	}

	if course == nil {
		return nil, status.Errorf(codes.NotFound, "GetCourse: course not found")
	}

	grpcCourse, err := models.BindCourseToGRPC(course)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetCourse: could not bind course to grpc: %v", err)
	}

	return &pb.GetCourseResponse{
		Course: grpcCourse,
	}, nil
}
