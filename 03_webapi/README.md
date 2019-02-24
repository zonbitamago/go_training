# webapi

本ディレクトリでは、Go言語を利用してWebAPIを作成する。

参考：[https://qiita.com/yukpiz/items/64e2874fd37a9dc7365d](https://qiita.com/yukpiz/items/64e2874fd37a9dc7365d)

ホットリロード：[https://github.com/codegangsta/gin](https://github.com/codegangsta/gin)

## サーバ立ち上げ

Go言語標準パッケージの `"net/http"`を利用する。

[https://godoc.org/net/http](https://godoc.org/net/http)

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Go API!!")
	})

	port := ":8080"

	fmt.Println("start!", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}
```