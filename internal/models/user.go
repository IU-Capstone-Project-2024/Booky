package models

import (
	pb "booky-back/api/booky"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID        string               `json:"id"`
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Password  Password             `json:"password"`
	CreatedAt *timestamp.Timestamp `json:"created_at"`
}

func BindUser(grpcUser *pb.User) (*User, error) {
	if grpcUser == nil {
		return nil, fmt.Errorf("grpc user is nil")
	}

	password, err := BindPassword(grpcUser.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to bind password: %w", err)
	}

	return &User{
		ID:        grpcUser.Id,
		Name:      grpcUser.Name,
		Email:     grpcUser.Email,
		Password:  *password,
		CreatedAt: grpcUser.CreatedAt,
	}, nil
}

func BindUserToGRPC(user *User) (*pb.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	password, err := BindPasswordToGRPC(&user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to bind password to grpc: %w", err)
	}

	return &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  password,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *User) Validate() bool {
	return u.ID != "" &&
		u.Name != "" &&
		u.Email != "" &&
		u.Password.Validate() // Modify this line according to your Password struct
}
