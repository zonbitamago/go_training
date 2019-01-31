# interfaces_are_implemented_implicitly

[https://go-tour-jp.appspot.com/methods/10](https://go-tour-jp.appspot.com/methods/10)

型にメソッドを定義することによって、インターフェースを実装したことになる。インターフェースを実装することを明示的に宣言する必要はない("implements"キーワードは必要ない)。

```go
// I インターフェース
type I interface {
	M()
}

// T 構造体
type T struct {
	S string
}

// M インターフェース実装
// implementsキーワードは必要ない
func (t T) M() {
	fmt.Println(t.S)
}
```