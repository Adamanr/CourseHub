package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func NewConnection(host, username, password string) (*pgx.Conn, error) {
	connUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/postgres?sslmode=disable", username, password, host)

	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		slog.Error("Failed to connect to postgres", err)
		return nil, err
	}

	return conn, nil
}

func CloseConnection(conn *pgx.Conn) error {
	err := conn.Close(context.Background())
	if err != nil {
		slog.Error("Failed to close postgres connection", err)
		return err
	}

	return nil
}
