package main

import "fmt"

// I interface
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// interfaceを実装しないとランタイムエラーとなる
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
