package main

// import "fmt"
// import "math"

// 上記の記述も可。以下のようにまとめることが多い。
// まとめる記述方法が良いとされている。
import (
	"fmt"
	"math"
)

func main() {
	// math.Sqrt(n Int)は引数の平方根を返す。
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
