package main

import "fmt"

// Vertex 頂点
type Vertex struct {
	X, Y int
}

var (
	// 全て初期値を設定する。
	v1 = Vertex{1, 2}
	// Xのみ初期化。Yはintの初期値(0)となる。
	v2 = Vertex{X: 1}
	// 初期化設定なし。X,Yはintの初期値(0)となる。
	v3 = Vertex{}
	// Vertexのポインタ宣言の際も初期値設定が可能。
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
