# choosing_a_value_or_pointer

[https://go-tour-jp.appspot.com/methods/8](https://go-tour-jp.appspot.com/methods/8)

ポインタレシーバを使う理由は2つある。

1つ目は、レシーバが指す先の変数の値を変更するため。

2つ目は、メソッドの呼び出しごとに変数のコピーを避けるため。特に、レシーバが大きな構造体である場合に効率的である。

一般的には変数レシーバ、または、ポインタレシーバのどちらかで全てのメソッドを統一して定義し、混在させるべきではない。