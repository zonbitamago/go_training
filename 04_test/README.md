# test

参考

- [https://budougumi0617.github.io/2018/08/19/go-testing2018/](https://budougumi0617.github.io/2018/08/19/go-testing2018/)
- [https://swet.dena.com/entry/2018/01/16/211035](https://swet.dena.com/entry/2018/01/16/211035)
- [https://swet.dena.com/entry/2018/01/22/120155](https://swet.dena.com/entry/2018/01/22/120155)
- [https://swet.dena.com/entry/2018/01/29/141707](https://swet.dena.com/entry/2018/01/29/141707)

## testingパッケージ

[https://golang.org/pkg/testing/](https://golang.org/pkg/testing/)

Goでテストを作成するときは、testingパッケージを利用する。

テストは`xxx_test.go`というファイルを作成する。この命名規則に従うことで、`go build`のときは無視され、`go test`のときのみビルドされる。

テストメソッドは`*testing.T`を引数として、`TestXxx`という命名規則で作成する。`Test`のあとの単語もキャメルケースで命名する必要がある。`TestSum`は実行されるが`Testsum`は実行されない。

Goのテストは、基本的に`t.Fatalf`を利用して、失敗した場合のメッセージを記述する。

`hoge/example_test.go`

```go
func TestExampleSuccess(t *testing.T) {
	result, err := example("hoge")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}

func TestExampleFailed(t *testing.T) {
	result, err := example("huga")
	if err == nil {
		t.Fatal("failed test")
	}
	if result != 0 {
		t.Fatal("failed test")
	}
}
```

## テスト起動コマンド

テスト起動はテストファイルが存在するディレクトリで以下のコマンドで可能となる。

```sh
go test
```

```sh
PASS
ok      go_training/04_test     0.012s
```

詳細な結果を見たい場合`-v`オプションを加える。

```sh
go test -v
```

```sh
=== RUN   TestExampleSuccess
--- PASS: TestExampleSuccess (0.00s)
=== RUN   TestExampleFailed
--- PASS: TestExampleFailed (0.00s)
PASS
ok      go_training/04_test     0.012s
```

ディレクトリを指定してテストを実施することも出来る。

```sh
go test ./hoge
```

すべてのテストを実施したい場合は以下のコマンドとなる。

```sh
go test ./...
```

## アサーション

Goは公式ではアサーションを提供していない。理由は以下のリンクに記載されている通り。

[http://go.shibu.jp/faq.html#id21](http://go.shibu.jp/faq.html#id21)

ただし、アサーションを使いたい場合は以下のライブラリが利用可能。

[https://github.com/stretchr/testify](https://github.com/stretchr/testify)

## サブテスト(テスト階層化)

テストの階層化は`t.Run`を利用して実現可能となる。

`calc/calc_test.go`

```go
func TestSum(t *testing.T) {
	t.Run("sum_plus", func(t *testing.T) {
		want := 3
		if result := Sum(1, 2); result != want {
			t.Fatalf("want:%v ,actual:%v", want, result)
		}
	})

	t.Run("sum_minus", func(t *testing.T) {
		want := -3
		if result := Sum(-1, -2); result != want {
			t.Fatalf("want:%v ,actual:%v", want, result)
		}

	})
}
```

テストを実行すると以下の様になる。

```sh
go test -v ./calc
```

```sh
=== RUN   TestSum
=== RUN   TestSum/sum_plus
=== RUN   TestSum/sum_minus
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/sum_plus (0.00s)
    --- PASS: TestSum/sum_minus (0.00s)
PASS
ok      go_training/04_test/calc
```

## カバレッジ

`--cover`オプションを付けることで、カバレッジ情報を付与できる。

```sh
go test --cover ./hoge
```

```sh
ok      go_training/04_test/hoge        0.012s  coverage: 100.0% of statements
```

`-coverprofile ファイル名`でカバレッジレポートを出力できる。

```sh
go test -coverprofile cover.out ./...
```

```sh
cat cover.out
mode: set
go_training/04_test/hoge/example.go:5.40,6.20 1 1
go_training/04_test/hoge/example.go:9.2,9.43 1 1
go_training/04_test/hoge/example.go:6.20,8.3 1 1
go_training/04_test/calc/calc.go:4.28,6.2 1 1
```

上記で出力したカバレッジレポートをhtml形式に変換する。

```sh
go tool cover -html=cover.out -o cover.html
```

## 共通処理

### BeforeAll / AfterAll

`BeforeAll`や`AfterAll`を実現したい場合は

```go
func TestMain(m *testing.M)
```

を利用する。

```go
func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}
```

```sh
$ go test -v ./calc/
before all...
=== RUN   TestSum
=== RUN   TestSum/sum_plus
=== RUN   TestSum/sum_minus
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/sum_plus (0.00s)
    --- PASS: TestSum/sum_minus (0.00s)
=== RUN   TestSubtract
=== RUN   TestSubtract/subtract_plus
=== RUN   TestSubtract/subtract_minus
--- PASS: TestSubtract (0.00s)
    --- PASS: TestSubtract/subtract_plus (0.00s)
    --- PASS: TestSubtract/subtract_minus (0.00s)
PASS
after all...
ok      go_training/04_test/calc        0.017s
```

### BeforeEach / AfterEach

BeforeEach/AfterEach用の関数は存在しないので、各テスト内で共通処理を行う必要がある。

## エラー出力

エラー出力に関しては、以下のように関数化しておくと便利。

`testutils/testutils.go`

```go
package testutils

import "testing"

// ErrorfHandler エラーメッセージ出力処理
func ErrorfHandler(tb testing.TB, want interface{}, got interface{}) {
	tb.Helper()
	tb.Errorf("want = %d, got = %d", want, got)
}
```

## Table Driven Test (テーブル駆動テスト)

Goではテーブルを使ってテストを書くことが推奨される。
サブテストもテーブルに含めて実施するとテストの見通しが良くなる。

※ここでいう「テーブル」とは、構造体のことを指す(構造体がテーブルのように見えるため)。

テーブルを使ってテストを記述すると、後にバグが発生するパターンのテストケースを追加するだけで済む。

`calc/calc_test.go`

```go
func TestMultiple(t *testing.T) {
	test := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"plus_pluse", 2, 3, 6},
		{"plus_minus", 3, -4, -12},
		{"minus_minus", -4, -5, 20},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple(tt.a, tt.b); got != tt.want {
				testutils.ErrorfHandler(t, tt.want, got)
			}
		})
	}
}
```

## モック

Go言語でのモックは[gomock](https://github.com/golang/mock)を利用する。

gomockではインターフェイスに対してテスト用モック処理を注入する形となる。

以下に、gomockの使い方を記載する。

テスト対象：`book/bookinfo.go`

```go
package book

import "fmt"

// Bookinfo 書籍情報取得
func Bookinfo(book Book) string {
    //book.ShowISBNをモック化したい。
	return fmt.Sprintf("ISBN is %v", book.ShowISBN())
}
```

モック化対象：`book/book.go`

```go
package book

// Book 本
type Book interface {
	ShowISBN() string
}

// Novel 小説
type Novel struct {
	ISBN string
}

// ShowISBN ISBN情報取得
func (novel *Novel) ShowISBN() string {
	return novel.ISBN
}
```

### モックファイルの作成

まずは[gomock](https://github.com/golang/mock)に記載されている通り、以下のインストールを行う。

```sh
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

次に、以下のコマンドでモックファイルの作成を行う。

```sh
mockgen -source=book/book.go --destination book/mock_book/mock_book.go
```

### テストの作成

テスト：`book/bookinfo_test.go`

```go
func TestBookinfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

    // モックの変数化
    book := mock_book.NewMockBook(ctrl)
    // モック処理の戻り値を設定。
	book.EXPECT().ShowISBN().Return("mockedISBN")

	test := []struct {
		name string
		a    Book
		want string
	}{
		{"mock_sample", book, "ISBN is mockedISBN"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bookinfo(tt.a); got != tt.want {
				testutils.ErrorfHandler(t, tt.want, got)
			}
		})
	}
}
```

テスト結果

```sh
=== RUN   TestBookinfo
=== RUN   TestBookinfo/mock_sample
--- PASS: TestBookinfo (0.00s)
    --- PASS: TestBookinfo/mock_sample (0.00s)
PASS
```