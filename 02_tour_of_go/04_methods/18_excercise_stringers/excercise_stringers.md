# excercise_stringers

[https://go-tour-jp.appspot.com/methods/18](https://go-tour-jp.appspot.com/methods/18)

`IPAddr`型を実装する。IPアドレスをドットで4つに区切った表現で出力するため、`fmt.Stringer`インターフェースを実装すること。

例えば、`IPAddr{1, 2 , 3, 4}`は`"1.2.3.4"`と出力すること。

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```