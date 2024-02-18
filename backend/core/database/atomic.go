package database

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresStore struct {
	Db *sql.DB
}

type AtomicAction func(any) error

// Executes atomic transaction
func (store *PostgresStore) ExecTx(
	ctx context.Context,
	fn func(*sql.Tx) (any, error),
	after AtomicAction,
) error {
	tx, err := store.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	obj, err := fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	if after != nil {
		if err := after(obj); err != nil {
			return fmt.Errorf("after action failure. error: %v", err)
		}
	}

	return tx.Commit()
}
