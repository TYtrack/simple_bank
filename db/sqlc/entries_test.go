/*
 * @Author: your name
 * @Date: 2021-12-15 00:03:04
 * @LastEditTime: 2021-12-21 23:00:25
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/bank_project/db/sqlc/entries_test.go
 */

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	listArgs := CreateEntryParams{
		AccountID: 1,
		Amount:    43,
	}
	_, err := testQueries.CreateEntry(context.Background(), listArgs)
	require.NoError(t, err)
}

func TestListEntriesById(t *testing.T) {
	listArgs := ListEntriesByIdParams{
		AccountID: 1,
		Limit:     10,
		Offset:    0,
	}
	zz, err := testQueries.ListEntriesById(context.Background(), listArgs)
	require.NoError(t, err)
	t.Errorf("zzz: %v", zz)
}

func TestGetEntry(t *testing.T) {

	zz, err := testQueries.GetEntry(context.Background(), 3)
	require.NoError(t, err)
	t.Errorf("zzz: %v", zz)
}

func TestDeleteEntry(t *testing.T) {

	err := testQueries.DeleteEntry(context.Background(), 3)
	require.NoError(t, err)
}
