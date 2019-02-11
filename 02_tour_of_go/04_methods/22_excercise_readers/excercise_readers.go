package main

import "golang.org/x/tour/reader"

// MyReader 独自Reader構造体
type MyReader struct{}

func (a MyReader) Read(rb []byte) (n int, e error) {
	for n, e = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
