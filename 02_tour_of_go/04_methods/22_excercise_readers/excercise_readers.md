# excercise_readers

[https://go-tour-jp.appspot.com/methods/22](https://go-tour-jp.appspot.com/methods/22)

ASCII文字`'A'`の無限ストリームを出力する`Reader`型を実装すること。

```go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
```
