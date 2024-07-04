package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"booky-back/internal/validator"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	courseData, err := models.BindCourse(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "CreateCourse: could not bind course data: %v", err)
	}

	v := validator.New()
	if details, err := v.ValidateCourse(courseData); err != nil {
		return nil, status.Error(codes.InvalidArgument, "CreateCourse: validation error: invalid course data")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "CreateCourse: validation error: %v", details)
	}

	course, err := s.Storage.CreateCourse(courseData)
	if err != nil {
		return nil, status.Error(codes.Internal, "CreateCourse: could not create course")
	}

	grpcCourse, err := models.BindCourseToGRPC(course)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateCourse: could not bind course to grpc: %v", err)
	}

	return &pb.CreateCourseResponse{
		Course: grpcCourse,
	}, nil
}
