# webapi

本ディレクトリでは、Go言語を利用してWebAPIを作成する。

参考：[http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/](http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/)

ホットリロード：[https://github.com/codegangsta/gin](https://github.com/codegangsta/gin)

## サーバ立ち上げ

```sh
gin -i run main.go
```

## webサーバテンプレートコード

Go言語標準パッケージの `"net/http"`を利用する。

[https://godoc.org/net/http](https://godoc.org/net/http)

`main.go`

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

## ルーティング処理追加

[https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)を利用する。

`main.go`

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/", HelloIndex)

	port := ":8080"

	fmt.Println("start!", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, router))
}

// HelloIndex 疎通確認用
func HelloIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello Go API!!")
}
```

## レスポンスをJsonに変更

`main.go`

```go
// HelloIndex 疎通確認レスポンス用json構造体
type HelloIndexResponse struct {
	Messaage string `json:"message"`
}

// Index 疎通確認用
func HelloIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexResponse := HelloIndexResponse{Messaage: "Hello Go API!!"}
	json.NewEncoder(w).Encode(indexResponse)
}
```

## パッケージ分割

`main.go`

```go
package main

import (
	"fmt"
	"go_training/03_webapi/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/", handlers.HelloIndex)

	port := ":8080"

	fmt.Println("start!", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, router))
}
```

`handlers/helloindex.go`

```go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HelloIndexResponse 疎通確認レスポンス用json構造体
type HelloIndexResponse struct {
	Messaage string `json:"message"`
}

// HelloIndex 疎通確認用
func HelloIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexResponse := HelloIndexResponse{Messaage: "Hello Go API!!"}
	json.NewEncoder(w).Encode(indexResponse)
}
```

## ログ出力追加

`logger/logger.go`

```go
package logger

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var logger = func(method, uri, name string, start time.Time) {
	log.Printf("\"method\":%q  \"uri\":%q    \"name\":%q   \"time\":%q", method, uri, name, time.Since(start))
}

// Logging ログ出力処理
func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}
```

`main.go`

```go
package main

import (
	"fmt"
	"go_training/03_webapi/handlers"
	"go_training/03_webapi/logger"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/", logger.Logging(handlers.HelloIndex, "index"))

	port := ":8080"

	fmt.Println("start!", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, router))
}
```

## レスポンスヘッダ修正

`handlers/helloindex.go`

```go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HelloIndexResponse 疎通確認レスポンス用json構造体
type HelloIndexResponse struct {
	Messaage string `json:"message"`
}

// HelloIndex 疎通確認用
func HelloIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexResponse := HelloIndexResponse{Messaage: "Hello Go API!!"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(indexResponse); err != nil {
		panic(err)
	}
}
```

