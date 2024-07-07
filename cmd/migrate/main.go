package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Naumovets/go-auth/internal/db/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg, err := postgres.NewConfig(".env")

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(1)
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.USER,
		cfg.PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)
	m, err := migrate.New("file://tools/migrate/migrations", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Applied migration: %d, Dirty: %t\n", version, dirty)

	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal(err)
	// }
}
