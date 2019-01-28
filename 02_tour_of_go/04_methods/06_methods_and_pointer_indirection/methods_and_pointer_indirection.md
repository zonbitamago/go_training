# methods_and_pointer_indirection

[https://go-tour-jp.appspot.com/methods/6](https://go-tour-jp.appspot.com/methods/6)

ポインタを引数に取る`ScaleFunc`関数はポインタを渡す必要がある。

```go
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

メソッドがポインタレシーバである場合は呼び出し時に変数、ポインタのいずれもレシーバとして受け取ることができる。

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

