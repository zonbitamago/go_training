# switch_evaluation_order

[https://go-tour-jp.appspot.com/flowcontrol/10](https://go-tour-jp.appspot.com/flowcontrol/10)

switch文は上から下へcaseを評価する。caseの条件が一致すれば、そこで停止(自動的にbreak)する。

例)

```go
swicth i {
    case 0:
    case f():
}
```

`i==0`であれば、`case 0`でbreakされるため`f`は呼び出されない。
