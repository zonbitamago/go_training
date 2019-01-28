# methods_continued

[https://go-tour-jp.appspot.com/methods/3](https://go-tour-jp.appspot.com/methods/3)

structの型以外にも、任意の型(`type`)にもメソッドを宣言することができる。

以下は`MyFloat`型に`Abs`メソッドを定義している。

```go
// MyFloat 内部float型
type MyFloat float64

// Abs 絶対値返却
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
```

レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージ内に存在する必要がある。他のパッケージで定義している型に対して、レシーバを伴うメソッドを宣言することはできない。