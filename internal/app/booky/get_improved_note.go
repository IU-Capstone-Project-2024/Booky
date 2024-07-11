package booky

import (
	pb "booky-back/api/booky"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetImprovedNote(ctx context.Context, req *pb.GetImprovedNoteRequest) (*pb.GetImprovedNoteResponse, error) {
	result, err := s.GPT.GetImprovedNote(req.GetBody())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting improved note: %v", err)
	}

	return &pb.GetImprovedNoteResponse{
		Body:         req.GetBody(),
		ImprovedBody: result}, nil
}
