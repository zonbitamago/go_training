# mutating_maps

[https://go-tour-jp.appspot.com/moretypes/22](https://go-tour-jp.appspot.com/moretypes/22)

マップの要素の追加、更新

```go
m[key] = elem
```

要素の取得

```go
elem = m[key]
```

要素の削除

```go
delete(m, kye)
```

キーに対する要素が存在するかどうかは、2つ目の返却値で判断する

```go
elem ,ok = m[key]
```

`m`に値があれば`ok`は`true`となり、存在しなければ`ok`は`false`となる。

mapに`key`が存在しない場合は、`elem`はmaptの要素の型のゼロ値(初期値)となる。
