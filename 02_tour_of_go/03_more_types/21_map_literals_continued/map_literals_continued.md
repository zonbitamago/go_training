# map_literals_continued

[https://go-tour-jp.appspot.com/moretypes/21](https://go-tour-jp.appspot.com/moretypes/21)

もしmapに渡す型がリテラルの要素から推定できる場合は、省略することができる。

```go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```
