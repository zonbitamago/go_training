package main

import "fmt"

// 型定義は定義の後ろに記載する。
// func func名(引数 型、引数 型) 戻り値型
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
