package postgres

import (
	"fmt"

	"github.com/Naumovets/go-auth/internal/db"
	"github.com/go-pg/pg"
)

func NewConn(cfg db.Config) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DB_HOST, cfg.DB_PORT),
		User:     cfg.USER,
		Password: cfg.PASSWORD,
		Database: cfg.DB_NAME,
	})

	return db, nil
}
