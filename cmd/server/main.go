package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/Naumovets/go-auth/internal/auth"
	"github.com/Naumovets/go-auth/internal/db/postgres"
	"github.com/Naumovets/go-auth/internal/repositories"
	desc "github.com/Naumovets/go-auth/pkg/auth_v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	httpAddress = "0.0.0.0:8080"
	grpcAddress = "0.0.0.0:50051"
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

	Cfg, err := auth.NewConfig(".auth.env")

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(1)
	}

	rep := repositories.NewRepository(conn)

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(2)
	}

	ctx := context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := startGrpcServer(&Cfg, rep); err != nil {
			log.Fatalf("gRPC error: %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := startHttpServer(ctx); err != nil {
			log.Fatalf("http error: %s", err)

		}
	}()

	wg.Wait()
}

func startGrpcServer(cfg *auth.Config, rep *repositories.Repository) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	desc.RegisterAuthV1Server(grpcServer, auth.NewServerAuth(rep, cfg))

	lis, err := net.Listen("tcp", grpcAddress)

	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at: %s\n", grpcAddress)

	return grpcServer.Serve(lis)
}

func startHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterAuthV1HandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(mux)

	log.Printf("http server listening at: %s\n", httpAddress)

	return http.ListenAndServe(httpAddress, handler)
}
