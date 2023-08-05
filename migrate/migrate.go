package main

import (
	"fmt"
	"todo/db"
	"todo/model"
)

func main() {
	// dbディレクトリのNewDB関数, 戻り値はDBインスタンスのアドレス
	dbConn := db.NewDB()
	defer fmt.Println("Successflly Migrated")
	defer db.CloseDB(dbConn)
	// dbに反映させたいモデル構造定義を引数に
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
