# functions

https://go-tour-jp.appspot.com/basics/5

関数の型は以下のように定義する。  
***func 関数名(引数 型, 引数 型) 戻り値型 { 処理 }***

```go
func add(x int, y int) int {
	return x + y
}
```

変数名の後ろに型名をつける形式にした理由については[公式ブログ](https://blog.golang.org/gos-declaration-syntax)に詳細な理由が記載してある。
