package main

import "fmt"

func main() {
	// 呼び出し元の関数(この場合はmain)の最後に遅延実行される。
	defer fmt.Println("world")

	fmt.Print("hello")
}
