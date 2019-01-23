# excercise_slices

[https://go-tour-jp.appspot.com/moretypes/18](https://go-tour-jp.appspot.com/moretypes/18)

画像生成のための`Pic`関数を実装すること。

この関数は、長さ`dy`のsliceの中に、長さ`dx`のsliceをを含む形で実装すること。

一番内部のsliceに好きな値を設定してください。面白い関数に`(x+y)/2`、`x*y`、`x^y`などがあります。

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    // この中身を実装すること。
}

func main() {
	pic.Show(Pic)
}
```