/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-29 13:37:11
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-29 20:22:58
 * @FilePath: /bank_project/db/sqlc/user_test.go
 */

package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashPassword, user2.HashPassword)
	require.Equal(t, user1.Username, user2.Username)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)

}
