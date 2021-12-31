/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-29 19:22:41
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-29 19:28:12
 * @FilePath: /bank_project/util/password_test.go
 */

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	pwd1 := RandomString(6)
	hashed_pwd_1, err := HashPassword(pwd1)
	require.NoError(t, err)

	err2 := CheckPassword(hashed_pwd_1, pwd1)
	require.NoError(t, err2)

	hashed_pwd_2, err := HashPassword(pwd1)
	require.NoError(t, err)
	require.NotEqual(t, hashed_pwd_1, hashed_pwd_2)

	pwd2 := RandomString(6)
	err = CheckPassword(hashed_pwd_1, pwd2)
	require.Equal(t, err.Error(), bcrypt.ErrMismatchedHashAndPassword.Error())

}
