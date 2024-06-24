package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"booky-back/internal/storage"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	course, err := models.BindCourse(req.GetCourse())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "UpdateCourse: could not bind course: %v", err)
	}

	updatedCourse, err := s.Storage.UpdateCourse(course)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "UpdateCourse: course not found")
		}
		return nil, status.Errorf(codes.Internal, "UpdateCourse: could not update course: %v", err)
	}

	grpcCourse, err := models.BindCourseToGRPC(updatedCourse)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "UpdateCourse: could not bind course to grpc: %v", err)
	}

	return &pb.UpdateCourseResponse{
		Course: grpcCourse,
	}, nil
}
