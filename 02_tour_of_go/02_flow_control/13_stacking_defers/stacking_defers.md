# stacking_defers

[https://go-tour-jp.appspot.com/flowcontrol/13](https://go-tour-jp.appspot.com/flowcontrol/13)

deferステートメントへ複数の関数を渡した場合は、呼び出し元の関数がreturnするときにLIFO(Last-In-First-Out)の順番で実行される。

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
```

実行結果は以下の通りとなる。

```text
counting
done
9
8
7
6
5
4
3
2
1
0
```