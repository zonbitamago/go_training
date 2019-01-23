# range_continued

[https://go-tour-jp.appspot.com/moretypes/17](https://go-tour-jp.appspot.com/moretypes/17)

インデックスや値は"`_`"(アンダーバー)へ代入することで捨てることができる。

```go
for _, value := range pow {
    fmt.Printf("%d\n", value)
}
```

インデックスのみが必要な場合は"`, value`"を省略する。

```go
for i := range pow {
    pow[i] = 1 << uint(i) // == 2**i
}
```
