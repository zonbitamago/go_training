# excercise_loops_and_funcgions

[https://go-tour-jp.appspot.com/flowcontrol/8](https://go-tour-jp.appspot.com/flowcontrol/8)

## 演習問題

平方根の近似値を計算する関数を作成しなさい。

### ヒント(バビロニア人の方法)

[https://cpplover.blogspot.com/2010/11/blog-post_20.html](https://cpplover.blogspot.com/2010/11/blog-post_20.html)

以下の計算繰り返していくと、平方根の値に近づいていくことができる。

```Text
x1 = (x0 + S/x0) /2
x2 = (x1 + S/x1) /2
```

例)√123の近似値を求める場合。

```Text
x1 = ( x0 + 123 / x0 ) / 2 = ( 1 + 123 / 1 ) / 2 = 62
x2 = ( 62 + 123 / 62 ) / 2 = 31.9919
x3 = 17.9183
x4 = 12.3914
x5 = 11.1588
x6 = 11.0905
```
