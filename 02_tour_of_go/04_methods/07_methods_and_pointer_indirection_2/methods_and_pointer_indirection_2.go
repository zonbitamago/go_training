package main

import (
	"fmt"
	"math"
)

// Vertex 頂点
type Vertex struct {
	X, Y float64
}

// Abs 斜辺
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AbsFunc 斜辺(関数)
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
}
