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

type ExecTransfersReq struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	Amount        int64 `json:"amount"`
}

type ExecTransfersResult struct {
	Transfer    Transfer `json:"transfer"`
	ToAccount   Account  `json:"to_account"`
	FromAccount Account  `json:"from_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *Store) ExecTransfers(ctx context.Context, arg ExecTransfersReq) (ExecTransfersResult, error) {
	var transferResult ExecTransfersResult

	err := s.execTransaction(ctx, func(q *Queries) error {
		var err error
		transferResult.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		transferResult.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		transferResult.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// TODO: Update the account balance

		return nil

	})

	return transferResult, err
}
