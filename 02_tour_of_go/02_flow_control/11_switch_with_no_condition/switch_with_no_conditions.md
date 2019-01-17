# switch_with_no_conditions

[https://go-tour-jp.appspot.com/flowcontrol/11](https://go-tour-jp.appspot.com/flowcontrol/11)

条件の無いswitchは`switch true`と記載することと同じとなる。

```go
switch /* true */ {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```