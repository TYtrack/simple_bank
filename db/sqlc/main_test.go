/*
 * @Author: your name
 * @Date: 2021-12-14 18:21:53
 * @LastEditTime: 2021-12-15 14:00:00
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/银行项目/db/sqlc/main_tets.go
 */

package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://zplus:123456@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println(err)
		return
	}

	testQueries = New(testDB)

	//运行单元测试
	os.Exit(m.Run())
}
