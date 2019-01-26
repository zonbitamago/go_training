# excercise_maps

[https://go-tour-jp.appspot.com/moretypes/23](https://go-tour-jp.appspot.com/moretypes/23)

`WordCount`関数を実装すること。

`WordCount`関数は、引数`s`で渡される文章の、各単語の出現回数のmapを返却する。

```go
package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    // ここを実装すること。
	return map[string]int{"x": 1}
}

func main() {
    // wc.Testは本処理のテスト関数
	wc.Test(WordCount)
}
```