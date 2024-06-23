package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) UpdateCourse(context.Context, *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	return &pb.UpdateCourseResponse{}, nil
}
