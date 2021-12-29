/*
 * @Author: your name
 * @Date: 2021-12-15 00:03:04
 * @LastEditTime: 2021-12-29 14:04:55
 * @LastEditors: TYtrack
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /bank_project/db/sqlc/entries_test.go
 */

package db

import (
	"bank_project/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func createRandomEntry(t *testing.T) (Account, Entry) {
	account := createRandomAccount(t)
	listArgs := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt64(10, 100),
	}
	entry, err := testQueries.CreateEntry(context.Background(), listArgs)
	require.NoError(t, err)
	return account, entry

}

func TestListEntriesById(t *testing.T) {
	account, _ := createRandomEntry(t)

	listArgs := ListEntriesByIdParams{
		AccountID: account.ID,
		Limit:     10,
		Offset:    0,
	}
	_, err := testQueries.ListEntriesById(context.Background(), listArgs)
	require.NoError(t, err)
}

func TestGetEntry(t *testing.T) {
	_, entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)

	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)

}

func TestDeleteEntry(t *testing.T) {
	_, entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
}
