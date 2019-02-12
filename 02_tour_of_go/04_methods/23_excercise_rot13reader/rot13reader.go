package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {

	n, err = reader.r.Read(b)

	if err != nil {
		return 0, err
	}

	for i := range b {
		// 各文字に ROT13 を適用する
		b[i] = rot13(b[i])
	}

	return
}

func rot13(c byte) byte {
	// 文字を ROT13 変換する関数
	switch {
	case ('A' <= c && c <= 'Z'):
		// 13 文字分ずらす
		return (c-'A'+13)%26 + 'A'
	case ('a' <= c && c <= 'z'):
		// 13 文字分ずらす
		return (c-'a'+13)%26 + 'a'
	default:
		// 何もしない
		return c
	}
}

func main() {
	s := strings.NewReader("Lbh penpexrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
