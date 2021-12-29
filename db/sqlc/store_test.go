/*
 * @Author: your name
 * @Date: 2021-12-15 13:15:06
 * @LastEditTime: 2021-12-29 16:09:39
 * @LastEditors: TYtrack
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /bank_project/db/sqlc/store_test.go
 */

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	account_1 := createRandomAccount(t)
	account_2 := createRandomAccount(t)

	store := NewStore(testDB)
	amount := 1000
	args := TransferTxParams{
		From_account_id: account_1.ID,
		To_account_id:   account_2.ID,
		Amount:          int64(amount),
	}

	n := 5
	errChan := make(chan error)
	resChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {

		go func() {
			res, err := store.TransferTx(context.Background(), args)
			if err != nil {
				t.Logf("zzz %v", err)
			}
			errChan <- err
			resChan <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan
		res := <-resChan
		require.NoError(t, err)
		require.NotEmpty(t, res)

		transfer := res.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account_1.ID, transfer.FromAccountID)
		require.Equal(t, account_2.ID, transfer.ToAccountID)
		require.Equal(t, transfer.Amount, args.Amount)
		require.NotZero(t, transfer.CreatedAt)
		require.NotZero(t, transfer.ID)

		fromEntry := res.From_entry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account_1.ID, fromEntry.AccountID)
		require.Equal(t, -fromEntry.Amount, args.Amount)
		require.NotZero(t, fromEntry.CreatedAt)
		require.NotZero(t, fromEntry.ID)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := res.To_entry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account_2.ID, toEntry.AccountID)
		require.Equal(t, toEntry.Amount, args.Amount)
		require.NotZero(t, toEntry.CreatedAt)
		require.NotZero(t, toEntry.ID)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		fromAccount := res.From_account
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account_1.ID, fromAccount.ID)

		toAccount := res.To_account
		require.NotEmpty(t, toAccount)
		require.Equal(t, account_2.ID, toAccount.ID)

		diff1 := account_1.Balance - fromAccount.Balance
		diff2 := -account_2.Balance + toAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%int64(amount) == 0)

	}

}

func TestTransferTxDeadlock(t *testing.T) {
	account_1 := createRandomAccount(t)
	account_2 := createRandomAccount(t)

	store := NewStore(testDB)
	amount := 1000

	n := 10
	errChan := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account_1.ID
		toAccountID := account_2.ID

		if i%2 == 0 {
			fromAccountID = account_2.ID
			toAccountID = account_1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				From_account_id: fromAccountID,
				To_account_id:   toAccountID,
				Amount:          int64(amount),
			})

			errChan <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan

		require.NoError(t, err)
	}
	updateAccount_1, err := testQueries.GetAccountForUpdate(context.Background(), account_1.ID)
	require.NoError(t, err)

	updateAccount_2, err := testQueries.GetAccountForUpdate(context.Background(), account_2.ID)
	require.NoError(t, err)

	require.Equal(t, updateAccount_1.Balance, account_1.Balance)
	require.Equal(t, updateAccount_2.Balance, account_2.Balance)

}
