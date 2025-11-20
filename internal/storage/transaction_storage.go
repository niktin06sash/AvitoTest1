package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TxManagerImpl struct {
	db *DBObject
}

func NewTxManager(db *DBObject) *TxManagerImpl {
	return &TxManagerImpl{
		db: db,
	}
}
func (txm *TxManagerImpl) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return txm.db.pool.BeginTx(ctx, pgx.TxOptions{})
}
func (txm *TxManagerImpl) Commit(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}
func (txm *TxManagerImpl) Rollback(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}
