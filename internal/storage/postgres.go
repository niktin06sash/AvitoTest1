package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBObject struct {
	pool *pgxpool.Pool
}

func NewDBObject(connectionString string) (*DBObject, error) {
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.TODO(), poolConfig)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.TODO())
	if err != nil {
		pool.Close()
		return nil, err
	}
	log.Println("Successful connect to Postgres-Client")
	return &DBObject{pool: pool}, nil
}
func (db *DBObject) Close() {
	db.pool.Close()
	log.Println("Successful close Postgre-Client")
}
func (db *DBObject) Ping(ctx context.Context) error {
	err := db.pool.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}
