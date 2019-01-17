# defer

[https://go-tour-jp.appspot.com/flowcontrol/12](https://go-tour-jp.appspot.com/flowcontrol/12)

deferステートメントは`defer`へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる。

```go
func main() {
	// 呼び出し元の関数(この場合はmain)の最後に遅延実行される。
	defer fmt.Println("world")

	fmt.Print("hello")
}
```

deferステートメントへ渡した関数の引数はすぐに評価されるが、その関数自体は遅延実行となる。