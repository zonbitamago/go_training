# if

[https://go-tour-jp.appspot.com/flowcontrol/5](https://go-tour-jp.appspot.com/flowcontrol/5)

Go言語のifステートメントは括弧「()」は不要。中括弧「{}」は必要。  
(forステートメントと同じ。)

```go
// 平方根を返却
func sqrt(x float64) string {
	// 引数がマイナスの場合は虚数扱いの"i"を付与。
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
```