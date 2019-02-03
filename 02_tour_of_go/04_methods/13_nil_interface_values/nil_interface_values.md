# nil_interface_values

[https://go-tour-jp.appspot.com/methods/13](https://go-tour-jp.appspot.com/methods/13)

インターフェースを実装しない場合は、値も具体的な型も保持しない。

インターフェースを実装しないで実行した場合は、ランタイムエラーとなる。

```go
// I interface
type I interface {
	M()
}

func main() {
	var i I
	// interfaceを実装しないとランタイムエラーとなる
	i.M()
}
```
