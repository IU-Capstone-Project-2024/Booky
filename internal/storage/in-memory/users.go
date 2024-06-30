package inmemory

import (
	"booky-back/internal/models"

	"booky-back/internal/storage"
	"fmt"
)

func (s *InMemoryStorage) GetUser(id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return user, nil
}
