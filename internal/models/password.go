package models

import (
	pb "booky-back/api/booky"
	"fmt"
)

type Password struct {
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
}

func BindPassword(grpcPassword *pb.Password) (*Password, error) {
	if grpcPassword == nil {
		return nil, fmt.Errorf("grpc password is nil")
	}

	return &Password{
		Password:     grpcPassword.Password,
		PasswordHash: grpcPassword.PasswordHash,
	}, nil
}

func BindPasswordToGRPC(password *Password) (*pb.Password, error) {
	if password == nil {
		return nil, fmt.Errorf("password is nil")
	}

	return &pb.Password{
		Password:     password.Password,
		PasswordHash: password.PasswordHash,
	}, nil
}

func (p *Password) Validate() bool {
	return p.Password != "" && p.PasswordHash != ""
}
