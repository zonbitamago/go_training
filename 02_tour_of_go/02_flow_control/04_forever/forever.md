# forever

[https://go-tour-jp.appspot.com/flowcontrol/4](https://go-tour-jp.appspot.com/flowcontrol/4)

ループ条件を省略すると無限ループになってしまうので注意。

```go
func main() {
	// 実行すると無限ループになるので、実行しないこと。
	for {

	}
}
```