package main

import (
	"booky-back/internal/pkg/auth"
	"booky-back/internal/pkg/auth/basic_auth"
	"context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (app *Application) authorize(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {
	if isRequired := auth.IsAuthRequired(ctx, info.FullMethod); !isRequired {
		return ctx, nil
	}

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	basicAuth := basic_auth.New(app.Server.Storage)
	ctx, err := basicAuth.Authorize(ctx, meta)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return ctx, nil
}
