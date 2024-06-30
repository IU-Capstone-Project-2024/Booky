package models

import (
	pb "booky-back/api/booky"
	"fmt"
)

type Course struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description *string     `json:"description"`
	Tracks      []pb.Track  `json:"tracks"`
	Semester    pb.Semester `json:"semester"`
	Year        int         `json:"year"`
}

func BindCourse(courseData *pb.CreateCourseData) (*Course, error) {
	if courseData == nil {
		return nil, fmt.Errorf("grpc course is nil")
	}

	return &Course{
		Title:       courseData.Title,
		Description: courseData.Description,
		Tracks:      courseData.Tracks,
		Semester:    courseData.Semester,
		Year:        int(courseData.Year),
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
		Tracks:      course.Tracks,
		Semester:    course.Semester,
		Year:        int32(course.Year),
	}, nil
}

func (c *Course) Validate() bool {
	return c.Title != ""
}
