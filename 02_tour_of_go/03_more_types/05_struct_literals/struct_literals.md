# struct_literals

[https://go-tour-jp.appspot.com/moretypes/5](https://go-tour-jp.appspot.com/moretypes/5)

structリテラルはフィールドの値を列挙し初期値の割当を示す。

`Name:`構文を利用してフィールドを一部のみ初期値割当することができる。

```go
var (
	// 全て初期値を設定する。
	v1 = Vertex{1, 2}
	// Xのみ初期化。Yはintの初期値(0)となる。
	v2 = Vertex{X: 1}
	// 初期化設定なし。X,Yはintの初期値(0)となる。
	v3 = Vertex{}
	// Vertexのポインタ宣言の際も初期値設定が可能。
	p = &Vertex{1, 2}
)
```