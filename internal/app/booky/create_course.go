package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) CreateCourse(context.Context, *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	return &pb.CreateCourseResponse{}, nil
}
