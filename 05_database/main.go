package main

import (
	"fmt"
	"go_training/05_database/gormsample"
	"go_training/05_database/standard"
)

func main() {
	fmt.Println("===standard start!===")
	standard.Select()
	standard.Insert()
	standard.InsertWithError()
	fmt.Println("===standard end!===")
	fmt.Println("===gorm start!===")
	gormsample.Select()
	fmt.Println("===gorm end!===")
}
