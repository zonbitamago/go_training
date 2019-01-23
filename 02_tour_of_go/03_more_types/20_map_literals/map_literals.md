# map_literals

[https://go-tour-jp.appspot.com/moretypes/20](https://go-tour-jp.appspot.com/moretypes/20)

mapリテラルは、キー(key)が必要になる。

```go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.86433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```
