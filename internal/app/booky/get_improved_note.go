package booky

import (
	pb "booky-back/api/booky"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetImprovedNote(ctx context.Context, req *pb.GetImprovedNoteRequest) (*pb.GetImprovedNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImprovedNote not implemented")
}
