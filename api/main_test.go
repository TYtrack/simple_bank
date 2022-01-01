/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-28 00:02:31
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-01 16:02:41
 * @FilePath: /bank_project/api/main_test.go
 */

package api

import (
	db "bank_project/db/sqlc"
	"bank_project/util"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	symetric_key := util.RandomString(32)
	config := util.Config{
		TokenSymmetricKey:   symetric_key,
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
