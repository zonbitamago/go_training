package main

import (
	"fmt"
	"math"
)

// Vertex 頂点
type Vertex struct {
	X, Y float64
}

// Abs 絶対値返却
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
