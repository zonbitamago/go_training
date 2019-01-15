package main

import (
	"fmt"
	"math"
)

// 平方根を返却
func sqrt(x float64) string {
	// 引数がマイナスの場合は虚数扱いの"i"を付与。
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
