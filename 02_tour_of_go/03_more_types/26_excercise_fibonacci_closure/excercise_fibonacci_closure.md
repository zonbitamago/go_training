# excercise_fibonacci_closure

[https://go-tour-jp.appspot.com/moretypes/26](https://go-tour-jp.appspot.com/moretypes/26)

以下のフィボナッチ関数を実装すること。

```go
// フィボナッチ数を返却する関数
func fibonacci() func() int {
    // ここを実装すること。
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

結果は以下の通りとなる。

```txt
0
1
1
2
3
5
8
13
21
34
```