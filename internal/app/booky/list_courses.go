package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) ListCourses(context.Context, *pb.ListCoursesRequest) (*pb.ListCoursesResponse, error) {
	return &pb.ListCoursesResponse{}, nil
}
