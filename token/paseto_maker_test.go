/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-31 16:00:53
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-31 16:01:57
 * @FilePath: /bank_project/token/paseto_maker_test.go
 */

package token

import (
	"bank_project/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	secretKey := util.RandomString(32)
	maker, err := NewPasetoMaker(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := util.RandomString(6)
	duration := time.Second
	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, payload.Username, username)

	time.Sleep(time.Second)

	payload2, err2 := maker.VerifyToken(token)
	require.ErrorIs(t, err2, ErrExpiredToken)
	require.Empty(t, payload2)

}
