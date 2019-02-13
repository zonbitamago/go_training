# images

[https://go-tour-jp.appspot.com/methods/24](https://go-tour-jp.appspot.com/methods/24)

`image`パッケージは、以下の`Image`インターフェースを定義している。

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

Note:`Bounds`メソッドの戻り値である`Rectangle`は、`image`パッケージの`image.Rectangle`に定義がある。

(詳細は[こちら](https://golang.org/pkg/image/#Image))

`color.Color`と`color.Model`はともにインターフェースだが、定義済みの`color.RGBA`と`color.RGBAModel`を使うことで、このインターフェースを無視できる。これらのインターフェースは、[images/Color](https://golang.org/pkg/image/color/)パッケージで定義されている。