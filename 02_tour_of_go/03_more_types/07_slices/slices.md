# slices

[https://go-tour-jp.appspot.com/moretypes/7](https://go-tour-jp.appspot.com/moretypes/7)

配列は固定中だが、スライスは可変長である。より柔軟な配列とみなすこともできる。実際にはスライスは配列よりも一般的である。

型`[]T`は型`T`のスライスを表す。

転んで区切られた2つのインデックスlowとhighの境界を指定し、スライスを形成する。

```go
a[low:high]
```

これは最初の要素を含み、最後の要素は含まない。

次の式は要素の1から3を含むスライスを示す。

```go
primes := [6]int{2, 3, 4, 7, 11, 13}
// {3,4,7}となる
var s []int = primes[1:4]
```
