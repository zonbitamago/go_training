package main

import (
	"fmt"
)

// Sqrt 平方根を返す
func Sqrt(root float64) float64 {
	initial := float64(1)
	x := babylonianMethod(initial, root)
	lastX := float64(0)

	for x != lastX {
		lastX = x
		x = babylonianMethod(x, root)
	}
	return x
}

// babylonianMethod 平方根近似値計算(バビロニア人の方法)
func babylonianMethod(x float64, y float64) float64 {
	return (x + y/x) / 2
}

func main() {
	fmt.Println(Sqrt(2))
	// for i := float64(1); i <= 10; i++ {
	// 	fmt.Println(i, Sqrt(i))
	// }
}
