# pointers

[https://go-tour-jp.appspot.com/moretypes/1](https://go-tour-jp.appspot.com/moretypes/1)

Goはポインタを扱い、ポインタは値のメモリアドレスを指す。

変数`T`のポインタは、`*T`型で、ゼロは`nil`です。

```go
var p *int
```

`&`オペレータはオペランド(operand)へのポインタを引き出す(参照渡し状態となる)。

```go
i := 42
p = &i
```

`*`オペレータはポインタの指す先の変数を示す。

```go
fmt.Println(*p) // ポインタpを通してiから値を読みだす
*p = 21         // ポインタpを通してiへ値を代入する
```

