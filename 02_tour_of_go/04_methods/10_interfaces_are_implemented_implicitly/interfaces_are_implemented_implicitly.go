package main

import "fmt"

// I インターフェース
type I interface {
	M()
}

// T 構造体
type T struct {
	S string
}

// M インターフェース実装
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
