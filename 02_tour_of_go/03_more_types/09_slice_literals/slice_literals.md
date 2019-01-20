# slice_literals

[https://go-tour-jp.appspot.com/moretypes/9](https://go-tour-jp.appspot.com/moretypes/9)

スライスのリテラルは長さのない配列のようなものである。

下記は配列リテラルである。

```go
[3]bool{true, true, false}
```

下記は、上記と同様の配列を作成し、それを参照するスライスである。

```go
[]bool{true, true, false}
```
