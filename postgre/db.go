package postgre

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	sqlx.DB
}

func NewDB() (*DB, error) {
	db, err := sqlx.Open("postgres", "user=dev dbname=dev sslmode=disable password=dev host=localhost")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &DB{*db}, nil
}
