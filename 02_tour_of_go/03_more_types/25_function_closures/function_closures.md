# function_closures

[https://go-tour-jp.appspot.com/moretypes/25](https://go-tour-jp.appspot.com/moretypes/25)

Go言語の関数なクロージャ(closure)である。

下記`adder`関数はクロージャを返している。

```go
// adder()が呼び出された時点しか「sum:=0」は実行されない。
// returnされている関数はsumの値は加算されていく。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
    // pos,negの中身は「func(x int) int」となる。
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
        // 「pos(i)」は「func(x int)int」が実行され、「sum:=0」は実行されない
		fmt.Println(pos(i), neg(-2*i))
	}
}
```