# named_return_values

[https://go-tour-jp.appspot.com/basics/7](https://go-tour-jp.appspot.com/basics/7)

Goでは関数の戻り値定義の際に、戻り値となる変数を定義することができる。  
名前をつけた戻り値の変数を使うと、returnステートメントには何も記載しないことが可能となる。これを"naked" returnと呼ぶ。

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```