package booky

import (
	pb "booky-back/api/booky"
	"context"
)

func (s *Server) HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: true}, nil
}
