# maps

[https://go-tour-jp.appspot.com/moretypes/19](https://go-tour-jp.appspot.com/moretypes/19)

`maps`はキーと値を関連付けする。

マップのゼロ値(初期値)は`nil`である。

`make`関数は初期化されたマップを返す。

```go
type Vertex struct {
	Lat, Long float64
}

// map[key型]Value型 の形でmapを定義する
var m map[string]Vertex

// mapの初期化された値を返す。
m = make(map[string]Vertex)
```
