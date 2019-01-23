package main

import (
	"golang.org/x/tour/pic"
)

// Pic 画像情報を設定
func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := range ret {
		ret[i] = make([]uint8, dx)
		for j := range ret[i] {
			ret[i][j] = uint8(i * j)
		}
	}
	return ret
}

func main() {
	// 画像データをバイト文字列形式で出力
	pic.Show(Pic)
}
