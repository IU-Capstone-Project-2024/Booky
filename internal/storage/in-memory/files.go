package inmemory

import (
	"booky-back/internal/models"
	"booky-back/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FileStorage struct {
	files map[string]*models.File
}

func NewFileStorage() *FileStorage {
	return &FileStorage{
		files: make(map[string]*models.File),
	}
}

func (s *FileStorage) CreateFile(file *models.File) (*models.File, error) {
	file.ID = uuid.New().String()
	file.CreatedAt = timestamppb.Now()

	s.files[file.ID] = file
	return file, nil
}

func (s *FileStorage) GetFile(id string) (*models.File, error) {
	file, ok := s.files[id]
	if !ok {
		return nil, fmt.Errorf("file with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return file, nil
}

func (s *FileStorage) DeleteFile(id string) error {
	_, ok := s.files[id]
	if !ok {
		return fmt.Errorf("file with id %s was not found: %w", id, storage.ErrNotFound)
	}

	delete(s.files, id)
	return nil
}

func (s *FileStorage) ListFiles(courseID string) ([]*models.File, error) {
	fileList := make([]*models.File, 0, len(s.files))
	for _, f := range s.files {
		if f.CourseID == courseID {
			fileList = append(fileList, f)
		}
	}
	return fileList, nil
}
