# slice_length_and_capacity

[https://go-tour-jp.appspot.com/moretypes/11](https://go-tour-jp.appspot.com/moretypes/11)

スライスは長さ(length)と容量(capacity)の性質を持つ。

スライスの長さは、それに含まれる要素数。

スライスの容量は、スライスの最初の要素から数えた、元の配列の要素数。

```go
s := []int{2, 3, 5, 7, 11, 13}

// lengthが減る
s = s[:0]

// lengthが増える
s = s[:4]

// capacityが減る
s = s[2:]
```

スライス`s`の長さは`len(s)`と`cap(s)`という式で得ることができる。