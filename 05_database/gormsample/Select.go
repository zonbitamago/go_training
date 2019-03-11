package gormsample

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Users usersテーブルDTO
type Users struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// Select select文発行
func Select() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}