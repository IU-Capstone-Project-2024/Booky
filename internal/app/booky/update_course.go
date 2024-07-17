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

func (s *Server) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
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
			return nil, status.Errorf(codes.NotFound, "UpdateCourse: course not found")
		}
		return nil, status.Errorf(codes.Internal, "UpdateCourse: could not get course: %v", err)
	}

	err = course.BindUpdateCourse(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "UpdateCourse: could not bind course: %v", err)
	}

	if details, err := v.ValidateCourse(course); err != nil {
		return nil, status.Error(codes.InvalidArgument, "UpdateCourse: validation error: invalid course data")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "UpdateCourse: validation error: %v", details)
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
