package standard

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User userテーブルdto
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// Select データ取得
func Select() {
	// 第2引数の形式は "user:password@tcp(host:port)/dbname"
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			id
			,name
		FROM users
		WHERE id = ?
	`, 1)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&(user.ID), &(user.Name))
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
}
