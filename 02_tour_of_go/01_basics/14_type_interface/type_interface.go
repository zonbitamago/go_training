package main

import "fmt"

func main() {
	// 明示的に型指定を行わず変数を宣言すると、型推論が行われる。

	// 以下はintとなる。
	// v := 42

	// 以下はstringとなる。
	v := "42"
	fmt.Printf("v is of type %T\n", v)
}
