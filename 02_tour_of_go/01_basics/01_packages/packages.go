package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intnは疑似乱数。
	// http://golang.jp/pkg/rand
	fmt.Println("My favorite number is", rand.Intn(10))
}
