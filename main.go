/*
 * @Author: your name
 * @Date: 2021-12-22 12:51:29
 * @LastEditTime: 2022-01-01 16:50:16
 * @LastEditors: TYtrack
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /bank_project/main.go
 */

package main

import (
	"bank_project/api"
	db "bank_project/db/sqlc"
	"bank_project/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	fmt.Println("config,", config)

	if err != nil {
		log.Fatalln("Error: config file cannot load", err)
	}

	db_conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println(err)
		return
	}

	store := db.NewStore(db_conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
