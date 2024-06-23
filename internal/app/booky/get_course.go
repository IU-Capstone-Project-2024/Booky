package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) GetCourse(context.Context, *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	return &pb.GetCourseResponse{}, nil
}
