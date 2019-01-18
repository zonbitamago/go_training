# structs

[https://go-tour-jp.appspot.com/moretypes/2](https://go-tour-jp.appspot.com/moretypes/2)

`struct`(構造体)はフィールドの集まりです。

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

実行結果

```txt
{1 2}
```