package main

import (
	"fmt"
	"math"
)

// MyFloat 内部float型
type MyFloat float64

// Abs 絶対値返却
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
