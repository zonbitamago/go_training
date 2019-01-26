package main

import "fmt"

// フィボナッチ数を返却する関数
func fibonacci() func() int {
	before2 := 0
	before1 := 0
	retVal := 0
	return func() int {
		retVal = before2 + before1
		before2 = before1
		before1 = retVal
		if before2 == 0 && before1 == 0 {
			before2 = 1
		}
		return retVal
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
