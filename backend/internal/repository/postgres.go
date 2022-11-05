package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func DBConnection(sugar *zap.SugaredLogger) (*pgxpool.Pool, error) {
	username := "postgres"
	password := "postgres"
	host := "172.20.0.4"
	port := 5432
	dbname := "tumaris"
	sslmode := "disable"
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&statement_cache_mode=describe", username, password, host, port, dbname, sslmode)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pool, err := pgxpool.Connect(ctx, dbURI)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		sugar.Errorf("Error occured while trying to ping in postgres: %v", err)
		return nil, err
	}

	return pool, nil
}
