# readers

[https://go-tour-jp.appspot.com/methods/21](https://go-tour-jp.appspot.com/methods/21)

`io`パッケージは、データストリームを読むことを表現する`io.Reader`インターフェースを規定している。

Goの標準ライブラリには、インターフェース、ファイル、ネットワーク接続、圧縮、暗号化、などで[多くの実装](https://golang.org/search?q=Read#Global)がある。

`io.Reader`インターフェースは`Read`メソッドを持つ。

```go
func (T) Read(b []byte) (n int, err error)
```

`Read`は、データを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返す。ストリームの終端は、`io.EOF`のエラーで返す。

例のコードは、`strings.Reader`を作成し、8バイト毎に読み込む。