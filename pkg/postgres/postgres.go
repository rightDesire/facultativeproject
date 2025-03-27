package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Creds описывает конфигурацию подключения к PostgreSQL.
// Обычно эти данные берутся из переменных окружения или из .env
type Creds struct {
	DSN string // DSN = Data Source Name, например: "postgres://user:pass@host:5432/dbname?sslmode=disable"
}

// NewGDB инициализирует пул соединений к PostgreSQL.
func NewGDB(ctx context.Context, creds Creds) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(creds.DSN)
	if err != nil {
		return nil, fmt.Errorf("parse config error: %w", err)
	}

	// Настраиваем параметры пула, если нужно
	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 2 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create pool error: %w", err)
	}

	// Проверяем, что подключение успешно
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping db error: %w", err)
	}

	return pool, nil
}
