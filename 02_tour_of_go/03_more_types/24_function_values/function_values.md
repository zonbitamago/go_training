# function_values

[https://go-tour-jp.appspot.com/moretypes/24](https://go-tour-jp.appspot.com/moretypes/24)

関数も変数として扱うことができる。

```go
// 変数の中身が関数となる。
hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(5, 12))
```

関数の引数に別の関数を取ることもできるし、戻り値としても利用できる。

```go
// 引数が「float64の引数2つ、戻り値がfloat64」の関数となる。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
```
