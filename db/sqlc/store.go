/*
 * @Author: your name
 * @Date: 2021-12-15 12:41:45
 * @LastEditTime: 2021-12-21 20:43:28
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/bank_project/db/sqlc/store.go
 */

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

var txKey = struct{}{}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) exeTx(ct context.Context, fnq func(*Queries) error) (err error) {
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	q := New(tx)
	err = fnq(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("err:%v  fer:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit()

}

type TransferTxParams struct {
	From_account_id int64 `json:"from_account_id"`
	To_account_id   int64 `json:"to_account_id"`
	Amount          int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer     Transfer `json:"transfer"`
	From_account Account  `json:"from_account"`
	To_account   Account  `json:"to_account"`
	From_entry   Entry    `json:"from_entry"`
	To_entry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.exeTx(ctx, func(q *Queries) error {

		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.From_account_id,
			ToAccountID:   arg.To_account_id,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.From_entry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.From_account_id,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.From_entry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.To_account_id,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		//TODO :资金的流入需要考虑死锁

		if arg.From_account_id < arg.To_account_id {
			result.From_account, result.To_account, err =
				transMoney(ctx, q, arg.From_account_id, -arg.Amount, arg.To_account_id, arg.Amount)

			if err != nil {
				return err
			}
		} else {
			result.To_account, result.From_account, err =
				transMoney(ctx, q, arg.To_account_id, arg.Amount, arg.From_account_id, -arg.Amount)

			if err != nil {
				return err
			}
		}
		return nil
	})
	return result, err
}

func transMoney(ctx context.Context,
	q *Queries,
	account_id_1 int64,
	amount_1 int64,
	account_id_2 int64,
	amount_2 int64,
) (account_1 Account, account_2 Account, err error) {
	account_1, err = q.AddAccountAmount(ctx, AddAccountAmountParams{
		ID:     account_id_1,
		Amount: amount_1,
	})
	if err != nil {
		return
	}
	account_2, err = q.AddAccountAmount(ctx, AddAccountAmountParams{
		ID:     account_id_2,
		Amount: amount_2,
	})
	if err != nil {
		return
	}
	return

}
