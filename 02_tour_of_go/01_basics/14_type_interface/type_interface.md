# type_interface

[https://go-tour-jp.appspot.com/basics/14](https://go-tour-jp.appspot.com/basics/14)

明示的な型を指定せずに変数を宣言する場合は、変数の型は右側の値から型推論される。

右側の変数が型を持っている場合は、左側の変数は右側の型と同じ型となる。

```go
var i int
j := i // j is an int
```

右側が変数ではなく値の場合は、値の方と同じ型となる。

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
s := "test"       // string
```