# imports
[https://go-tour-jp.appspot.com/basics/2](https://go-tour-jp.appspot.com/basics/2)

別パッケージのインポートは以下のようにimportを宣言して行う。

```go
import {
    "fmt"
    "math"
}
```

下記の記載方法も可能だが、前述のようにまとめて記載するほうが良いとされている。

```go
import "fmt"
import "math"
```