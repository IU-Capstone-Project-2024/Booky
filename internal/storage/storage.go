package storage

import "booky-back/internal/models"

type Storage struct {
	CourseStorage CourseModel
	NoteStorage   NoteModel
	FileStorage   FileModel
	UserStorage   UserModel
}

func (s *Storage) CreateCourse(course *models.Course) (*models.Course, error) {
	return s.CourseStorage.CreateCourse(course)
}

func (s *Storage) GetCourse(id string) (*models.Course, error) {
	return s.CourseStorage.GetCourse(id)
}

func (s *Storage) UpdateCourse(course *models.Course) (*models.Course, error) {
	return s.CourseStorage.UpdateCourse(course)
}

func (s *Storage) DeleteCourse(id string) error {
	return s.CourseStorage.DeleteCourse(id)
}

func (s *Storage) ListCourses() ([]*models.Course, error) {
	return s.CourseStorage.ListCourses()
}

func (s *Storage) CreateNote(note *models.Note) (*models.Note, error) {
	return s.NoteStorage.CreateNote(note)
}

func (s *Storage) GetNote(id string) (*models.Note, error) {
	return s.NoteStorage.GetNote(id)
}

func (s *Storage) UpdateNote(note *models.Note) (*models.Note, error) {
	return s.NoteStorage.UpdateNote(note)
}

func (s *Storage) DeleteNote(id string) error {
	return s.NoteStorage.DeleteNote(id)
}

func (s *Storage) ListNotes(courseID string) ([]*models.Note, error) {
	return s.NoteStorage.ListNotes(courseID)
}

func (s *Storage) CreateFile(file *models.File) (*models.File, error) {
	return s.FileStorage.CreateFile(file)
}

func (s *Storage) GetFile(id string) (*models.File, error) {
	return s.FileStorage.GetFile(id)
}

func (s *Storage) DeleteFile(id string) error {
	return s.FileStorage.DeleteFile(id)
}

func (s *Storage) ListFiles(courseID string) ([]*models.File, error) {
	return s.FileStorage.ListFiles(courseID)
}

func (s *Storage) GetUser(id string) (*models.User, error) {
	return s.UserStorage.GetUser(id)
}
