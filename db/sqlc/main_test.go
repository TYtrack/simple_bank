/*
 * @Author: your name
 * @Date: 2021-12-14 18:21:53
 * @LastEditTime: 2021-12-22 16:23:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/银行项目/db/sqlc/main_tets.go
 */

package db

import (
	"bank_project/util"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalln("Error: config file cannot load", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println(err)
		return
	}

	testQueries = New(testDB)

	//运行单元测试
	os.Exit(m.Run())
}
