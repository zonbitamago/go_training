# interface_values

[https://go-tour-jp.appspot.com/methods/11](https://go-tour-jp.appspot.com/methods/11)

インターフェースの値は、値と具体的な型の2つからなる。

```txt
(value, type)
```

インターフェースの値は、特定の基底になる具体的な型の値を保持する。

インターフェースの値のメソッドを呼び出すと、その基底型の同じ名前のメソッドが実行される。