package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListCourses(ctx context.Context, req *pb.ListCoursesRequest) (*pb.ListCoursesResponse, error) {
	courses, err := s.Storage.ListCourses()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ListCourses: failed to list courses: %v", err)
	}

	grpcCourses := make([]*pb.Course, 0, len(courses))
	for _, c := range courses {
		grpcCourse, err := models.BindCourseToGRPC(c)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "ListCourses: could not bind course to grpc: %v", err)
		}
		grpcCourses = append(grpcCourses, grpcCourse)
	}

	return &pb.ListCoursesResponse{
		Courses: grpcCourses,
	}, nil
}
