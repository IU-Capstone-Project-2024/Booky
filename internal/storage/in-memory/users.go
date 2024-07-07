package inmemory

import (
	"booky-back/internal/models"

	"booky-back/internal/storage"
	"fmt"

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
		ID:    "93d47db0-f1e2-401a-bc67-e564e3f26107",
		Name:  "Admin",
		Email: "admin@example.com",
		Password: models.Password{
			PasswordHash: "admin",
		},
		CreatedAt: timestamppb.Now(),
	}

	nikita := &models.User{
		ID:    "7cc3f064-d595-4676-904a-b3c31a3e3038",
		Name:  "Nikita Shlyakhtin",
		Email: "nikita@example.com",
		Password: models.Password{
			PasswordHash: "password1",
		},
		CreatedAt: timestamppb.Now(),
	}

	adel := &models.User{
		ID:    "1585483f-9e78-4aad-a626-bf160e1239ac",
		Name:  "Adel Shagaliev",
		Email: "adel@example.com",
		Password: models.Password{
			PasswordHash: "password2",
		},
		CreatedAt: timestamppb.Now(),
	}

	mikhail := &models.User{
		ID:    "1de67a82-24fc-4ea3-8618-c7d42bffc71d",
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
