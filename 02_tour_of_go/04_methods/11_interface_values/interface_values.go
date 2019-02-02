package main

import (
	"fmt"
	"math"
)

// I インターフェース
type I interface {
	M()
}

// T 構造体
type T struct {
	S string
}

// M インターフェース実装1
func (t *T) M() {
	fmt.Println(t.S)
}

// F float64
type F float64

// M インターフェース実装2
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
