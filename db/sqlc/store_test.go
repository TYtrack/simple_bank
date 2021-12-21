/*
 * @Author: your name
 * @Date: 2021-12-15 13:15:06
 * @LastEditTime: 2021-12-21 15:27:24
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/bank_project/db/sqlc/store_test.go
 */

package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	args := TransferTxParams{
		From_account_id: 8,
		To_account_id:   7,
		Amount:          1000,
	}

	n := 5
	errChan := make(chan error)
	resChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {

		go func() {
			res, err := store.TransferTx(context.Background(), args)
			if err != nil {
				t.Fatalf("zzz %v", err)
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

		fromAccount := res.From_account
		fmt.Printf("account :%v \t amount :%v\n", fromAccount.ID, fromAccount.Balance)

		toAccount := res.To_account
		fmt.Printf("account :%v \t amount :%v\n", toAccount.ID, toAccount.Balance)

	}

}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	n := 10
	errChan := make(chan error)
	resChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		var from_account int64 = 10
		var to_account int64 = 9
		if i%2 == 1 {
			from_account = 9
			to_account = 10
		}

		go func() {
			fmt.Printf("%v  --> %v\n", from_account, to_account)
			res, err := store.TransferTx(context.Background(), TransferTxParams{
				From_account_id: from_account,
				To_account_id:   to_account,
				Amount:          1000,
			})
			if err != nil {
				t.Fatalf("zzz2 %v", err)
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

		fromAccount := res.From_account
		fmt.Printf("---account :%v \t amount :%v\n", fromAccount.ID, fromAccount.Balance)

		toAccount := res.To_account
		fmt.Printf("***account :%v \t amount :%v\n", toAccount.ID, toAccount.Balance)

	}

}
