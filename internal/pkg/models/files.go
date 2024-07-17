package models

import (
	pb "booky-back/api/booky"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type File struct {
	ID        string               `json:"id"`
	CourseID  string               `json:"course_id"`
	Content   []byte               `json:"content"`
	Filename  string               `json:"filename"`
	Publisher User                 `json:"publisher"`
	CreatedAt *timestamp.Timestamp `json:"created_at"`
}

func BindFile(fileData *pb.CreateFileData) (*File, error) {
	if fileData == nil {
		return nil, fmt.Errorf("grpc file is nil")
	}

	return &File{
		CourseID:  fileData.CourseId,
		Content:   fileData.Content,
		Filename:  fileData.Filename,
		Publisher: User{ID: fileData.UserId},
	}, nil
}

func BindFileToGRPC(file *File) (*pb.File, error) {
	if file == nil {
		return nil, fmt.Errorf("file is nil")
	}

	publisher, err := BindUserToGRPC(&file.Publisher)
	if err != nil {
		return nil, fmt.Errorf("failed to bind publisher to grpc: %w", err)
	}

	return &pb.File{
		Id:        file.ID,
		CourseId:  file.CourseID,
		Content:   file.Content,
		Filename:  file.Filename,
		Publisher: publisher,
		CreatedAt: file.CreatedAt,
	}, nil
}

func (f *File) Validate() bool {
	return f.CourseID != "" &&
		f.Filename != "" &&
		len(f.Content) > 0 &&
		f.Publisher.ID != ""
}
