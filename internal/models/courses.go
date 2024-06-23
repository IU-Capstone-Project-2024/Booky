package models

import (
	pb "booky-back/api/booky"
	"fmt"
)

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func BindCourse(grpcCourse *pb.Course) (*Course, error) {
	if grpcCourse == nil {
		return nil, fmt.Errorf("grpc course is nil")
	}

	return &Course{
		ID:          grpcCourse.Id,
		Title:       grpcCourse.Title,
		Description: grpcCourse.Description,
	}, nil
}

func BindCourseToGRPC(course *Course) (*pb.Course, error) {
	if course == nil {
		return nil, fmt.Errorf("course is nil")
	}

	return &pb.Course{
		Id:          course.ID,
		Title:       course.Title,
		Description: course.Description,
	}, nil
}

func (c *Course) Validate() bool {
	return c.Title != "" && c.Description != ""
}
