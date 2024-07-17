package auth

import (
	"context"
)

type Auth interface {
	Authorize(context.Context, map[string][]string) (context.Context, error)
}

// List of methods that do not require authentication
var noAuthMethods = map[string]bool{
	"/booky.BookyService/HealthCheck": true,
	"/booky.BookyService/CreateUser":  true,
}

func IsAuthRequired(ctx context.Context, method string) bool {
	if _, ok := noAuthMethods[method]; ok {
		return false
	}
	return true
}
