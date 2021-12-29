/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-28 00:02:31
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-28 00:02:31
 * @FilePath: /bank_project/api/main_test.go
 */

package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
