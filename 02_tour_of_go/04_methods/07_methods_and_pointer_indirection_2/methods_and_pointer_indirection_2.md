# methods_and_pointer_indirection_2

[https://go-tour-jp.appspot.com/methods/7](https://go-tour-jp.appspot.com/methods/7)

変数の引数を取る関数は、特定の型の変数を取る必要がある。

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

メソッドが変数レシーバである場合は呼び出し時に変数、または、ポインタのいずれかのレシーバをして取ることができる。

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
