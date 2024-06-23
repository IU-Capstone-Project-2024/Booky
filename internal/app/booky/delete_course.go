package booky

import (
	pb "booky-back/api/booky"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	err := s.Storage.DeleteCourse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "DeleteCourse: could not delete course: %v", err)
	}

	return &pb.DeleteCourseResponse{}, nil
}
