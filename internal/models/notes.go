package models

import (
	pb "booky-back/api/booky"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type Note struct {
	ID        string               `json:"id"`
	CourseID  string               `json:"course_id"`
	Title     string               `json:"title"`
	Body      string               `json:"body"`
	CreatedAt *timestamp.Timestamp `json:"created_at"`
	UpdatedAt *timestamp.Timestamp `json:"updated_at"`
	Publisher User                 `json:"publisher"`
}

func BindNote(noteData *pb.CreateNoteData) (*Note, error) {
	if noteData == nil {
		return nil, fmt.Errorf("grpc note is nil")
	}

	return &Note{
		CourseID:  noteData.CourseId,
		Title:     noteData.Title,
		Body:      noteData.Body,
		Publisher: User{ID: noteData.UserId},
	}, nil
}

func BindNoteToGRPC(note *Note) (*pb.Note, error) {
	if note == nil {
		return nil, fmt.Errorf("note is nil")
	}

	publisher, err := BindUserToGRPC(&note.Publisher)
	if err != nil {
		return nil, fmt.Errorf("failed to bind publisher to grpc: %w", err)
	}

	return &pb.Note{
		Id:        note.ID,
		CourseId:  note.CourseID,
		Title:     note.Title,
		Body:      note.Body,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
		Publisher: publisher,
	}, nil
}

func (n *Note) Validate() bool {
	return n.CourseID != "" && n.Title != "" && n.Body != ""
}
