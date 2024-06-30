package inmemory

import (
	"booky-back/internal/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type InMemoryStorage struct {
	courses map[string]*models.Course
	notes   map[string]*models.Note
	users   map[string]*models.User
}

func NewInMemoryStorage() *InMemoryStorage {
	user := &models.User{
		ID:    "1",
		Name:  "admin",
		Email: "admin@admin.admin",
		Password: models.Password{
			PasswordHash: "admin",
		},
		CreatedAt: timestamppb.Now(),
	}

	users := make(map[string]*models.User)
	users[user.ID] = user

	return &InMemoryStorage{
		courses: make(map[string]*models.Course),
		notes:   make(map[string]*models.Note),
		users:   users,
	}
}
