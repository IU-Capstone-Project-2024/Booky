package storage

import (
	models "booky-back/internal/pkg/models"
)

type DataModel interface {
	CourseModel
	NoteModel
	FileModel
	UserModel
}

type CourseModel interface {
	// Course CRUD
	CreateCourse(*models.Course) (*models.Course, error)
	GetCourse(string) (*models.Course, error)
	UpdateCourse(*models.Course) (*models.Course, error)
	DeleteCourse(string) error
	ListCourses() ([]*models.Course, error)
}

type NoteModel interface {
	// Note CRUD
	CreateNote(*models.Note) (*models.Note, error)
	GetNote(string) (*models.Note, error)
	UpdateNote(*models.Note) (*models.Note, error)
	DeleteNote(string) error
	ListNotes(string) ([]*models.Note, error)
}

type FileModel interface {
	// Files
	CreateFile(*models.File) (*models.File, error)
	GetFile(string) (*models.File, error)
	DeleteFile(string) error
	ListFiles(string) ([]*models.File, error)
}

type UserModel interface {
	// User CRUD
	CreateUser(*models.User) (*models.User, error)
	GetUser(string) (*models.User, error)
	DeleteUser(string) error
	ListUsers() ([]*models.User, error)
	FindUserByEmail(string) (*models.User, error)
}
