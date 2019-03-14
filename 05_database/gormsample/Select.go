package gormsample

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Users usersテーブルDTO
type Users struct {
	ID   int
	Name string
}

// Select select文発行
func Select() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
	if err != nil {
		panic(err.Error())
	}

	var users Users
	// 1件取得
	db.First(&users)
	fmt.Printf("db.First:%v\n", users)

	// 複数件取得
	var allUsers []Users
	db.Find(&allUsers)
	fmt.Printf("db.Find:%v\n", allUsers)

	defer db.Close()
}
