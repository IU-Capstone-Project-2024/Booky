package inmemory

import (
	"booky-back/internal/pkg/logger"
	"booky-back/internal/pkg/models"
	"booky-back/internal/pkg/storage"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
		ID:           "93d47db0-f1e2-401a-bc67-e564e3f26107",
		Name:         "Admin",
		Email:        "admin@example.com",
		PasswordHash: "$2a$10$x7O9tMHy83afMDfMW.koGORrzNE0Az76v9DYw1v/O1yd83t2kczMS", // password: password
		CreatedAt:    timestamppb.Now(),
	}

	nikita := &models.User{
		ID:           "7cc3f064-d595-4676-904a-b3c31a3e3038",
		Name:         "Nikita Shlyakhtin",
		Email:        "nikita@example.com",
		PasswordHash: "$2a$10$x7O9tMHy83afMDfMW.koGORrzNE0Az76v9DYw1v/O1yd83t2kczMS", // password: password
		CreatedAt:    timestamppb.Now(),
	}

	adel := &models.User{
		ID:           "1585483f-9e78-4aad-a626-bf160e1239ac",
		Name:         "Adel Shagaliev",
		Email:        "adel@example.com",
		PasswordHash: "$2a$10$x7O9tMHy83afMDfMW.koGORrzNE0Az76v9DYw1v/O1yd83t2kczMS", // password: password
		CreatedAt:    timestamppb.Now(),
	}

	mikhail := &models.User{
		ID:           "1de67a82-24fc-4ea3-8618-c7d42bffc71d",
		Name:         "Mikhail Zimin",
		Email:        "mikhail@example.com",
		PasswordHash: "$2a$10$x7O9tMHy83afMDfMW.koGORrzNE0Az76v9DYw1v/O1yd83t2kczMS", // password: password
		CreatedAt:    timestamppb.Now(),
	}

	return []*models.User{admin, nikita, adel, mikhail}
}

func (s *UserStorage) CreateUser(user *models.User) (*models.User, error) {
	if _, exists := s.users[user.ID]; exists {
		return nil, fmt.Errorf("user with id %s already exists: %w", user.ID, storage.ErrAlreadyExists)
	}

	user.ID = uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}
	user.PasswordHash = string(hashedPassword)

	logger.Infof("Created user with hash: %s", user.PasswordHash)

	user.CreatedAt = timestamppb.Now()
	user.UpdatedAt = timestamppb.Now()

	s.users[user.ID] = user

	return user, nil
}

func (s *UserStorage) GetUser(id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return user, nil
}

func (s *UserStorage) DeleteUser(id string) error {
	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user with id %s was not found: %w", id, storage.ErrNotFound)
	}

	delete(s.users, id)

	return nil
}

func (s *UserStorage) ListUsers() ([]*models.User, error) {
	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users, nil
}

func (s *UserStorage) FindUserByEmail(email string) (*models.User, error) {
	for _, user := range s.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with email %s was not found: %w", email, storage.ErrNotFound)
}
