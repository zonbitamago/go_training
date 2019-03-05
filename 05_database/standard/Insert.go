package standard

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Insert 登録処理
func Insert() {
	// 第2引数の形式は "user:password@tcp(host:port)/dbname"
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	// recoverでロールバックするように設定しておく
	// recoverについては以下。
	// https://blog.amedama.jp/entry/2015/10/11/123535
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			if err := tx.Rollback(); err != nil {
				panic(err.Error())
			}
			fmt.Println("Rollbacked")
		}
	}()

	_, err = tx.Exec(`
		INSERT INTO users(name) VALUES(?)
	`, "dummy")
	if err != nil {
		panic(err.Error())
	}

	// コミット
	if err = tx.Commit(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Commited")
}

// InsertWithError 登録処理(ロールバックされる)
func InsertWithError() {
	// 第2引数の形式は "user:password@tcp(host:port)/dbname"
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	// recoverでロールバックするように設定しておく
	// recoverについては以下。
	// https://blog.amedama.jp/entry/2015/10/11/123535
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			if err := tx.Rollback(); err != nil {
				panic(err.Error())
			}
			fmt.Println("Rollbacked")
		}
	}()

	_, err = tx.Exec(`
		INSERT INTO users(dummy) VALUES(?)
	`, "dummy")
	if err != nil {
		panic(err.Error())
	}

	// コミット
	if err = tx.Commit(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Commited")
}
