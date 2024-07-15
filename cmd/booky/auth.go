package main

import (
	"context"

	"booky-back/internal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// List of methods that do not require authentication
var noAuthMethods = map[string]bool{
	"/booky.BookyService/HealthCheck": true,
	"/booky.BookyService/CreateUser":  true,
}

func authorize(ctx context.Context, storage storage.Storage) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	var email, password string
	if val, ok := md["email"]; ok {
		email = val[0]
	} else {
		return nil, status.Error(codes.Unauthenticated, "missing email")
	}

	if val, ok := md["password"]; ok {
		password = val[0]
	} else {
		return nil, status.Error(codes.Unauthenticated, "missing password")
	}

	err := storage.VerifyUser(email, password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	return ctx, nil
}
