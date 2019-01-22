# creating_slice_with_make

[https://go-tour-jp.appspot.com/moretypes/13](https://go-tour-jp.appspot.com/moretypes/13)

スライスは`make`関数を使って作成することができる。`make`関数は動的サイズの配列を作成できる。

`make`関数はゼロ化された配列を割り当て、その配列を指すスライスを返す。

```go
a := make([]int, 5) // len(a) = 5
```

`make`関数の3番目の引数に、スライスの容量(capacity)を指定できる。

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```
