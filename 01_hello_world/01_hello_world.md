# 01_hello_world

## Goをインストールする。
[公式ページ](https://golang.org/)からインストーラをダウンロードし、インストールする。  
※その他のインストール方法は以下を参照すること。
[Getting Started(公式)](https://golang.org/doc/install)

## Goのワークスペースでディレクトリを作成する。
1.ワークスペースに以下のディレクトリを作成する。  
{ワークスペース}/src/01_hello_world  
※インストーラを使った場合は、ワークスペースは以下となる。
- Windowsの場合・・・```c:\Go```
- Macの場合・・・```~/go```

※ワークスペースを変更したい場合は以下を参照すること。  
[Getting Started(公式)](https://golang.org/doc/install)

## main.goを作成する。
上記ディレクトリに以下を配置する。

main.go
```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello go World!\n")
}
```

## ビルド&実行する。
以下のコマンドでビルドを行う。  
ビルドが完了すると、「01_hello_world」が作成される。

```
cd {ワークスペースのパス}/src/01_hello_world
go build
```

上記ファイルを実行する。

```
./01_hello_world
Hello go World!
```

※各々のIDEで実行する場合は、ビルドが不要となることがある。
IDEでの実行は各自で調べること。
