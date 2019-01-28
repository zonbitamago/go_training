package main

import "fmt"

// Vertex 頂点
type Vertex struct {
	X, Y float64
}

// Scale 倍数
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc 倍化関数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{3, 4}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
