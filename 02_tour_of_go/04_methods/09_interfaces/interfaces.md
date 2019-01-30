# interfaces

[https://go-tour-jp.appspot.com/methods/9](https://go-tour-jp.appspot.com/methods/9)

interface(インターフェース)型は、メソッドのシグニチャの集まりで定義する。

そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができる。

```go
// Abser 絶対値取得interface
type Abser interface {
	Abs() float64
}

// MyFloat float64拡張
type MyFloat float64

// Abs 絶対値取得
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())
}
```
