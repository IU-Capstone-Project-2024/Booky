package basic_auth

import (
	"booky-back/internal/pkg/storage"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type BasicAuth struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *BasicAuth {
	return &BasicAuth{
		storage: storage,
	}
}

func (a *BasicAuth) Authorize(ctx context.Context, meta map[string][]string) (context.Context, error) {
	var email, password string
	if val, ok := meta["email"]; ok {
		email = val[0]
	} else {
		return nil, fmt.Errorf("email address is required")
	}

	if val, ok := meta["password"]; ok {
		password = val[0]
	} else {
		return nil, fmt.Errorf("password is required")
	}

	user, err := a.storage.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	ctx = context.WithValue(ctx, "user", user)
	return ctx, nil
}
