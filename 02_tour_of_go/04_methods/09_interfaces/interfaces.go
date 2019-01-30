package main

import (
	"fmt"
	"math"
)

// Abser 絶対値取得interface
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	a = &v
	fmt.Println(a.Abs())

	// vは単純なVertexなので(*Vertexではないので)Abserのinterfaceにはなっておらず、エラーとなる。
	// a = v

}

// MyFloat float64拡張
type MyFloat float64

// Abs 絶対値取得
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex 頂点
type Vertex struct {
	X, Y float64
}

// Abs 斜辺
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
