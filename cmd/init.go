package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db *pgxpool.Pool
)

func connectDb(dsn string) error {
	var err error
	db, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	return nil
}
