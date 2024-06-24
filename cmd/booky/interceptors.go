package main

import (
	"context"
	"fmt"
	"strings"

	"booky-back/internal/logger"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func interceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		var fieldsStr string
		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]
			fieldsStr += fmt.Sprintf("%v=%v ", key, value)
		}

		fieldsStr = strings.TrimSpace(fieldsStr)

		finalMsg := fmt.Sprintf("%s %s", msg, fieldsStr)

		switch lvl {
		case logging.LevelDebug, logging.LevelInfo:
			logger.Info(finalMsg)
		case logging.LevelWarn:
			logger.Warn(finalMsg)
		case logging.LevelError:
			logger.Error(finalMsg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func unaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(interceptorLogger())
}

func unaryRecoveryInterceptor() grpc.UnaryServerInterceptor {
	handler := func(p any) (err error) {
		return status.Errorf(codes.Internal, "internal error: %v", p)
	}

	opts := []recovery.Option{
		recovery.WithRecoveryHandler(handler),
	}

	return recovery.UnaryServerInterceptor(opts...)
}
