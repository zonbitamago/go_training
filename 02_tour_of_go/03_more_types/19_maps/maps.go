package main

import "fmt"

// Vertex 頂点
type Vertex struct {
	Lat, Long float64
}

// map[key型]Value型 の形でmapを定義する
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
