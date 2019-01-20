# slices_are_like_references_to_arrays

[https://go-tour-jp.appspot.com/moretypes/8](https://go-tour-jp.appspot.com/moretypes/8)

スライスは配列への参照のようなものである。

スライスはデータを格納しておらず、元の配列の部分列を指し示している。

スライスの要素を変更すると、元となる配列の対応する要素が変更される。

同じ元の配列を共有している他のスライスは、その変更が反映される。