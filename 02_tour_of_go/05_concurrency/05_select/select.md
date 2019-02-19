# select

[https://go-tour-jp.appspot.com/concurrency/5](https://go-tour-jp.appspot.com/concurrency/5)

`select`ステートメントは、goroutineを複数の通信操作で待たせる。

`select`は、複数ある`case`のいずれかが準備できるようになるまでブロックし、準備が出来た`case`を実行する。もし、複数の`case`の準備ができている場合、`case`はランダムに選択される。