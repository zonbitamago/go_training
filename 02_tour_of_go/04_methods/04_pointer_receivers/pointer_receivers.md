# pointer_receivers

[https://go-tour-jp.appspot.com/methods/4](https://go-tour-jp.appspot.com/methods/4)

ポインタレシーバでメソッドを宣言することが可能である。

以下の例では`*Vertex`に`Scale`メソッドが定義されている。

```go
// 今までのレシーバ
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバ
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

ポインタレシーバを持つメソッドではレシーバが指す変数を変更することができる。(この場合だと、`v *Vertex`の中身を変更することができる。)

変数レシーバでは元の`Vertex`のコピーを操作する。そのため、引数はポインタレシーバで利用されることが多い。
