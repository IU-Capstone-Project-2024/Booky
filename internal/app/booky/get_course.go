package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/pkg/models"
	"booky-back/internal/pkg/storage"
	"booky-back/internal/pkg/validator"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	id := req.GetId()
	v := validator.New()
	if details, err := v.ValidateID(id); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid id")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	course, err := s.Storage.GetCourse(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "GetCourse: course not found")
		}
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
