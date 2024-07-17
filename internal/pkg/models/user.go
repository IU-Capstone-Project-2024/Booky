package models

import (
	pb "booky-back/api/booky"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID           string               `json:"id"`
	Name         string               `json:"name"`
	Email        string               `json:"email"`
	Password     string               `json:"-"`
	PasswordHash string               `json:"-"`
	CreatedAt    *timestamp.Timestamp `json:"created_at"`
	UpdatedAt    *timestamp.Timestamp `json:"updated_at"`
}

func BindUser(data *pb.CreateUserData) (*User, error) {
	if data == nil {
		return nil, fmt.Errorf("grpc user is nil")
	}

	return &User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}, nil
}

func BindUserToGRPC(user *User) (*pb.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	return &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *User) Validate() bool {
	return u.Name != "" &&
		u.Email != ""
}
