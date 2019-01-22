# slices_of_slices

[https://go-tour-jp.appspot.com/moretypes/14](https://go-tour-jp.appspot.com/moretypes/14)

スライスは他のスライスを含む任意の型を内部に持つことができる。

```go
board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}
```