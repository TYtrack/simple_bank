/*
 * @Author: your name
 * @Date: 2021-12-14 18:19:58
 * @LastEditTime: 2021-12-21 22:45:58
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/银行项目/db/sqlc/account_test.go
 */

package db

import (
	"context"
	"fmt"
	"testing"
	"util"

	"github.com/stretchr/testify/require"
)

func TestAccountCreate(t *testing.T) {
	accountParams := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), accountParams)
	if err != nil {
		fmt.Println(err)
	}
	//断言
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, accountParams.Owner, account.Owner)
	require.Equal(t, accountParams.Balance, account.Balance)
	require.Equal(t, accountParams.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func TestUpdateAccount(t *testing.T) {
	updateAccountParams := UpdateAccountParams{
		Balance: 899,
		ID:      2,
	}
	err := testQueries.UpdateAccount(context.Background(), updateAccountParams)
	require.NoError(t, err)
}

func TestUpdate2Account(t *testing.T) {
	updateAccoun2tParams := UpdateAccount2Params{
		Balance: 444,
		ID:      3,
	}
	account, err := testQueries.UpdateAccount2(context.Background(), updateAccoun2tParams)
	require.NoError(t, err)
	ss := fmt.Sprintf("zzz%v", account)
	t.Logf(ss)
}

func TestListAccounts(t *testing.T) {
	listAccountsParams := ListAccountsParams{
		Limit:  5,
		Offset: 1,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), listAccountsParams)
	require.NoError(t, err)
	t.Fatal(accounts)
}

func TestDeleteAccount(t *testing.T) {

	err := testQueries.DeleteAccount(context.Background(), 2)
	require.NoError(t, err)
}
