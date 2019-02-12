# rot13reader

[https://go-tour-jp.appspot.com/methods/23](https://go-tour-jp.appspot.com/methods/23)

`io.Reader`を実装し、全てのアルファベットに対してROT13変換を行う`rot13Reader`を実装すること。

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

