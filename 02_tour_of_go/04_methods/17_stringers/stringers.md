# stringers

[https://go-tour-jp.appspot.com/methods/17](https://go-tour-jp.appspot.com/methods/17)

最もよく使われているinterfaceの一つにfmtパッケージの`Stringer`がある。

```go
type Stringer interface {
    String() string
}
```

`Stringer`インターフェースはstringとして出力する際のフォーマットを指定するinterfaceである。