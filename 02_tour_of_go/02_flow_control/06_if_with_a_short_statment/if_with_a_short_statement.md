# if_with_a_short_statement

[https://go-tour-jp.appspot.com/flowcontrol/6](https://go-tour-jp.appspot.com/flowcontrol/6)

if文は条件判定時に変数宣言が可能。  
ここで宣言した変数はif文内のみ有効なスコープとなる。

```go
if v := math.Pow(x, n); v < lim {
    return v
}
```
