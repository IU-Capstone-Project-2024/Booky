package main

import (
	"booky-back/internal/app/booky"
	"booky-back/internal/config"
	"booky-back/internal/logger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "booky-back/api/booky"

	"google.golang.org/grpc"
)

type Application struct {
	Config *config.Config
	Server *booky.Server
}

func NewApp(config *config.Config) *Application {
	return &Application{
		Config: config,
		Server: booky.NewServer(),
	}
}

func (app *Application) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", app.Config.Ip, app.Config.Port))
	if err != nil {
		logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	server := booky.NewServer()
	pb.RegisterBookyServiceServer(grpcServer, server)

	done := make(chan bool)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		logger.InfoKV("shutting down server",
			"signal", s.String(),
		)

		grpcServer.GracefulStop()

		done <- true
	}()

	go func() {
		logger.InfoKV("starting server",
			"address", fmt.Sprintf("%s:%s", app.Config.Ip, app.Config.Port),
		)

		err = grpcServer.Serve(lis)
		if err != nil && err != grpc.ErrServerStopped {
			logger.Fatal(err)
		}
	}()

	<-done

	logger.Info("server gracefully stopped")

	return nil
}
