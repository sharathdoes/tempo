package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	config,err:=pgxpool.ParseConfig(databaseURL)
	if err!=nil {
		return nil, fmt.Errorf("parse database config : %w", err)
	}

	config.MaxConns=10
	config.MinConns=2

	ctx, cancel:=context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err:= pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}
	//Instead of creating a new database connection every time you execute a query, you maintain a collection of reusable connections.

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return pool, nil
}
