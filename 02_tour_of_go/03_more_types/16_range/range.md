# range

[https://go-tour-jp.appspot.com/moretypes/16](https://go-tour-jp.appspot.com/moretypes/16)

forループに利用するrangeは、スライスやマップ(`map`)を一つずつ反復処理するために使う。

スライスをrangeで繰り返す場合、rangeは反復ごとに2つの変数を返却する。1つ目はインデックスで、2つ目はインデックスの場所の要素のコピーである。

([公式ドキュメント](https://golang.org/ref/spec#RangeClause))