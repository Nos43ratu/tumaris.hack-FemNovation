package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func DBConnection(sugar *zap.SugaredLogger) (*pgxpool.Pool, error) {
	username := ""
	password := ""
	host := ""
	port := 5432
	dbname := ""
	sslmode := "disable"
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s&statement_cache_mode=describe", username, password, host, port, dbname, sslmode)
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
