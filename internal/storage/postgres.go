package storage

import (
	"AvitoTest1/config"
	"AvitoTest1/internal/logger"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBObject struct {
	pool *pgxpool.Pool
}

func NewDBObject(cfg config.DatabaseConfig) (*DBObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conurl := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	poolConfig, err := pgxpool.ParseConfig(conurl)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &DBObject{pool: pool}, nil
}
func (db *DBObject) Close(log *logger.Logger) {
	db.pool.Close()
	log.ZapLogger.Debug("Successful close database connect")
}
func (db *DBObject) Ping(ctx context.Context) error {
	err := db.pool.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}
