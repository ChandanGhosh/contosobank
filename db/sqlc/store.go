package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	transaction, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(transaction)
	err = fn(q)

	if err != nil {
		if rberr := transaction.Rollback(); rberr != nil {
			return fmt.Errorf("tx err: %v, rollback error: %v", err, rberr)
		}
	}

	return transaction.Commit()
}
