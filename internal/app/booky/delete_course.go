package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/storage"
	"booky-back/internal/validator"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	id := req.GetId()
	v := validator.New()
	if details, err := v.ValidateID(id); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid id")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	err := s.Storage.DeleteCourse(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "DeleteCourse: course not found")
		}
		return nil, status.Errorf(codes.Internal, "DeleteCourse: could not delete course: %v", err)
	}

	return &pb.DeleteCourseResponse{}, nil
}
