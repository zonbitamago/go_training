# if_and_else

[https://go-tour-jp.appspot.com/flowcontrol/7](https://go-tour-jp.appspot.com/flowcontrol/7)

ifステートメントで宣言された変数は、else内でも利用可能。

```go
func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		return lim
	}
}
```
