package main

import (
	"fmt"
	"math"
)

func main() {
	// 以下は、math.piが公開されていないため、エラーとなる。
	// fmt.Println(math.pi)

	// 以下は、math.Piが公開されているため、エラーとならない。
	fmt.Println(math.Pi)
}
