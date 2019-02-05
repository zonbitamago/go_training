# type_switches

[https://go-tour-jp.appspot.com/methods/16](https://go-tour-jp.appspot.com/methods/16)

型switchはいくつかの型アサーションを直列に使用できる。

型switchは通常のswitch文と似ていますが、型switchのcaseは型(値ではない)を指定し、それらの値は指定されたインターフェースの値が保持する値の型と比較される。

```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

型switchの宣言は、型アサーション`i.(T)`と同じ構文を持ちますが、特定の型`T`はキーワード`type`に置き換えられる。

このswitch文は、インターフェースの値`i`が型`T`または`S`の値を保持するかどうかをテストする。`T`および`S`の各caseにおいて、変数`v`はそれぞれ型`T`または`S`であり、`i`によって保持される値を保持する。defaultの場合(一致するものがない場合)、変数`v`は同じインターフェース型で値は`i`となる。

