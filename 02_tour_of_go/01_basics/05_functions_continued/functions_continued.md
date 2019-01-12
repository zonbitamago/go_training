# functions_continued

https://go-tour-jp.appspot.com/basics/5

関数の2つ以上の引数が同じ型の場合は、型定義をまとめて後ろに記載することができる。

(型名を個別に記載)

```go
func add(x int, y int) int {
	return x + y
}
```

(型名をまとめて記載)

```go
func add(x, y int) int {
	return x + y
}
```
