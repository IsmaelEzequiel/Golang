package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute the db queries and transactions
type Store struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// executes a function within a database transaction
func (store *Store) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rlErr := tx.Rollback(ctx); rlErr != nil {
			return fmt.Errorf(rlErr.Error())
		}
		return err
	}
	return tx.Commit(ctx)
}

type TransferForTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"Amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_accouunt"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// It creates a transfer record, add account entries, and update accounts balance within a single database transaction
func (store *Store) transferTx(ctx context.Context, arg TransferForTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTransaction(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update accounts
		return nil
	})

	return result, err
}
