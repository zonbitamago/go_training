# pointers_to_structs

[https://go-tour-jp.appspot.com/moretypes/4](https://go-tour-jp.appspot.com/moretypes/4)

structのフィールドはポインタを通してアクセスすることができる。

フィールド`X`を持つstructのポインタ`p`がある場合、フィールド`X`にアクセスするには`(*p).X`のように書くことができる。ただし、この書き方をすることは少なく、`p.X`と書くことが多い。
