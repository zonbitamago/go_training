# multiple_results

[https://go-tour-jp.appspot.com/basics/6](https://go-tour-jp.appspot.com/basics/6)

Go言語では関数の戻り値を複数定義することができる。

```go
// 複数戻り値を返却
func swap(x, y string) (string, string) {
	return y, x
}

// 複数の戻り値を受け取る
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```