# appending_to_a_slice

[https://go-tour-jp.appspot.com/moretypes/15](https://go-tour-jp.appspot.com/moretypes/15)

スライスへ要素を追加するには`append`関数を利用する。`append`関数は末尾に要素を追加する。

```go
func append(slice []Type, elems ...Type) []Type
```

(関数の詳細は[公式ドキュメント](https://golang.org/pkg/builtin/#append)を参照のこと。)

追加先の配列の容量が不足している場合は、より大きいサイズの配列を割り当て直し、戻り値のスライスは新しい割当先を示す。