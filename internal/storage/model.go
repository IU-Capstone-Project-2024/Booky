package storage

import "booky-back/internal/models"

type DataModel interface {
	// Course CRUD
	CreateCourse(*models.Course) (*models.Course, error)
	GetCourse(string) (*models.Course, error)
	UpdateCourse(*models.Course) (*models.Course, error)
	DeleteCourse(string) error
	ListCourses() ([]*models.Course, error)

	// Note CRUD
	CreateNote(*models.Note) (*models.Note, error)
	GetNote(string) (*models.Note, error)
	UpdateNote(*models.Note) (*models.Note, error)
	DeleteNote(string) error
	ListNotes(string) ([]*models.Note, error)

	// User CRUD
	GetUser(string) (*models.User, error)

	// Files
	CreateFile(*models.File) (*models.File, error)
	GetFile(string) (*models.File, error)
	DeleteFile(string) error
	ListFiles(string) ([]*models.File, error)
}
