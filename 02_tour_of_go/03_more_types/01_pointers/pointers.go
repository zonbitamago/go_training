package main

import "fmt"

func main() {
	i, j := 42, 2701

	// iのメモリアドレスがpに代入される(参照渡し状態)
	p := &i
	// ポインタの値を表示する。
	fmt.Println(*p)
	// ポインタの値を変更する。
	*p = 21
	// メモリアドレスを共有している他の変数でも値が変わっている。
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
