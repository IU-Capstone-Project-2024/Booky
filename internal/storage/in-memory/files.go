package inmemory

import (
	"booky-back/internal/models"
	"booky-back/internal/storage"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *InMemoryStorage) CreateFile(file *models.File) (*models.File, error) {
	file.ID = fmt.Sprint(len(s.files) + 1)
	file.CreatedAt = timestamppb.Now()

	s.files[file.ID] = file
	return file, nil
}

func (s *InMemoryStorage) GetFile(id string) (*models.File, error) {
	file, ok := s.files[id]
	if !ok {
		return nil, fmt.Errorf("file with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return file, nil
}

func (s *InMemoryStorage) DeleteFile(id string) error {
	_, ok := s.files[id]
	if !ok {
		return fmt.Errorf("file with id %s was not found: %w", id, storage.ErrNotFound)
	}

	delete(s.files, id)
	return nil
}

func (s *InMemoryStorage) ListFiles(courseID string) ([]*models.File, error) {
	fileList := make([]*models.File, 0, len(s.files))
	for _, f := range s.files {
		if f.CourseID == courseID {
			fileList = append(fileList, f)
		}
	}
	return fileList, nil
}
