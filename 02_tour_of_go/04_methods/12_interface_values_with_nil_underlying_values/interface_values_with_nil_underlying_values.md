# interface_values_with_nil_underlying_values

[https://go-tour-jp.appspot.com/methods/12](https://go-tour-jp.appspot.com/methods/12)

インターフェース自体の中にある具体的な値がnilの場合、メソッドはnilをレシーバとして呼び出される。

これは他の言語ではnullポインター例外を引き起こすことがありますが、Goはnilをレシーバとして呼び出されても適切に処理するメソッドを記述するのが一般的である。

```go
// M implements
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
```
