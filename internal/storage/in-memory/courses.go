package inmemory

import (
	"booky-back/internal/models"
	"booky-back/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

func (s *InMemoryStorage) CreateCourse(course *models.Course) (*models.Course, error) {
	course.ID = uuid.New().String()
	s.courses[course.ID] = course
	return course, nil
}

func (s *InMemoryStorage) GetCourse(id string) (*models.Course, error) {
	course, ok := s.courses[id]
	if !ok {
		return nil, fmt.Errorf("course with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return course, nil
}

func (s *InMemoryStorage) UpdateCourse(course *models.Course) (*models.Course, error) {
	_, ok := s.courses[course.ID]
	if !ok {
		return nil, fmt.Errorf("course with id %s was not found: %w", course.ID, storage.ErrNotFound)
	}

	s.courses[course.ID] = course
	return course, nil
}

func (s *InMemoryStorage) DeleteCourse(id string) error {
	_, ok := s.courses[id]
	if !ok {
		return fmt.Errorf("course with id %s was not found: %w", id, storage.ErrNotFound)
	}

	delete(s.courses, id)
	return nil
}

func (s *InMemoryStorage) ListCourses() ([]*models.Course, error) {
	courseList := make([]*models.Course, 0, len(s.courses))
	for _, c := range s.courses {
		courseList = append(courseList, c)
	}
	return courseList, nil
}
