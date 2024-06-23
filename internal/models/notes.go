package models

import (
	pb "booky-back/api/booky"
	"fmt"
)

type Note struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func BindNote(grpcNote *pb.Note) (*Note, error) {
	if grpcNote == nil {
		return nil, fmt.Errorf("grpc note is nil")
	}

	return &Note{
		ID:       grpcNote.Id,
		CourseID: grpcNote.CourseId,
		Title:    grpcNote.Title,
		Body:     grpcNote.Body,
	}, nil
}

func BindNoteToGRPC(note *Note) (*pb.Note, error) {
	if note == nil {
		return nil, fmt.Errorf("note is nil")
	}

	return &pb.Note{
		Id:       note.ID,
		CourseId: note.CourseID,
		Title:    note.Title,
		Body:     note.Body,
	}, nil
}

func (n *Note) Validate() bool {
	return n.CourseID != "" && n.Title != "" && n.Body != ""
}
