package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) DeleteCourse(context.Context, *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	return &pb.DeleteCourseResponse{}, nil
}
