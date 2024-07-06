package inmemory

import (
	"booky-back/internal/models"

	"booky-back/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserStorage struct {
	users map[string]*models.User
}

func NewUserStorage() *UserStorage {
	users := make(map[string]*models.User)

	testUsers := getTestUsers()
	for _, user := range testUsers {
		users[user.ID] = user
	}

	return &UserStorage{
		users: users,
	}
}

func getTestUsers() []*models.User {
	admin := &models.User{
		ID:    uuid.New().String(),
		Name:  "Admin",
		Email: "admin@example.com",
		Password: models.Password{
			PasswordHash: "admin",
		},
		CreatedAt: timestamppb.Now(),
	}

	nikita := &models.User{
		ID:    uuid.New().String(),
		Name:  "Nikita Shlyakhtin",
		Email: "nikita@example.com",
		Password: models.Password{
			PasswordHash: "password1",
		},
		CreatedAt: timestamppb.Now(),
	}

	adel := &models.User{
		ID:    uuid.New().String(),
		Name:  "Adel Shagaliev",
		Email: "adel@example.com",
		Password: models.Password{
			PasswordHash: "password2",
		},
		CreatedAt: timestamppb.Now(),
	}

	mikhail := &models.User{
		ID:    "4",
		Name:  "Mikhail Zimin",
		Email: "mikhail@example.com",
		Password: models.Password{
			PasswordHash: "password3",
		},
		CreatedAt: timestamppb.Now(),
	}

	return []*models.User{admin, nikita, adel, mikhail}
}

func (s *UserStorage) GetUser(id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return user, nil
}
