package db

import (
	"context"
	"log"

	"github.com/byterotom/infinity-play/config"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(ctx context.Context, env *config.Config) *pgx.Conn {
	conn, err := pgx.Connect(ctx, env.DatabaseUrl)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return conn
}
