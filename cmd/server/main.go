package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Naumovets/go-auth/internal/auth"
	"github.com/Naumovets/go-auth/internal/db/postgres"
	"github.com/Naumovets/go-auth/internal/repositories"
	desc "github.com/Naumovets/go-auth/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50051
)

func main() {

	pgCfg, err := postgres.NewConfig(".env")

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(1)
	}

	conn, err := postgres.NewConn(pgCfg)

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(1)
	}

	authCfg, err := auth.NewConfig(".auth.env")

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(1)
	}

	rep := repositories.NewRepository(conn)

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(2)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, auth.NewServerAuth(rep, &authCfg))

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
