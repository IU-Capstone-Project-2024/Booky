package inmemory

import (
	"booky-back/internal/models"
)

type InMemoryStorage struct {
	courses map[string]*models.Course
	notes   map[string]*models.Note
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		courses: make(map[string]*models.Course),
		notes:   make(map[string]*models.Note),
	}
}
