package main

import "fmt"

// Vertex 頂点の座標を示す構造体
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
