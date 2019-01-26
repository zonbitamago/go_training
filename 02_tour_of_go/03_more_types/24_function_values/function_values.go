package main

import (
	"fmt"
	"math"
)

// 引数が「float64の引数2つ、戻り値がfloat64」の関数となる。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 変数の中身が関数となる。
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
