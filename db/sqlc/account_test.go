/*
 * @Author: your name
 * @Date: 2021-12-14 18:19:58
 * @LastEditTime: 2022-01-01 20:35:01
 * @LastEditors: TYtrack
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /bank_project/db/sqlc/account_test.go
 */

package db

import (
	"bank_project/util"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountCreate(t *testing.T) {
	createRandomAccount(t)
}

func createRandomUser(t *testing.T) User {
	password := util.RandomString(6)
	hashPwd, _ := util.HashPassword(password)
	userParams := CreateUserParams{
		Username:     util.RandomString(6),
		HashPassword: hashPwd,
		FullName:     util.RandomString(10),
		Email:        util.RandomEmail(),
	}
	user1, err := testQueries.CreateUser(context.Background(), userParams)
	require.NoError(t, err)
	require.Equal(t, user1.FullName, userParams.FullName)
	require.Equal(t, user1.Email, userParams.Email)
	require.Equal(t, user1.HashPassword, userParams.HashPassword)
	require.Equal(t, user1.Username, userParams.Username)

	require.NotZero(t, user1.CreatedAt)
	require.True(t, user1.PasswordChangedAt.IsZero())
	return user1

}

func createRandomAccount(t *testing.T) Account {
	user1 := createRandomUser(t)
	accountParams := CreateAccountParams{
		Owner:    user1.Username,
		Balance:  10000,
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), accountParams)
	if err != nil {
		fmt.Println(err)
	}
	return account

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	fmt.Println("TestUpdateAccount")
	updateAccountParams := UpdateAccountParams{
		Balance: 899,
		ID:      account.ID,
	}
	err := testQueries.UpdateAccount(context.Background(), updateAccountParams)
	require.NoError(t, err)
}

func TestUpdate2Account(t *testing.T) {
	account := createRandomAccount(t)
	fmt.Println("TestUpdate2Account")
	updateAccoun2tParams := UpdateAccount2Params{
		Balance: 444,
		ID:      account.ID,
	}
	account, err := testQueries.UpdateAccount2(context.Background(), updateAccoun2tParams)
	require.NoError(t, err)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	fmt.Println("TestListAccounts")

	listAccountsParams := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), listAccountsParams)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	fmt.Println("TestDeleteAccount")
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}
