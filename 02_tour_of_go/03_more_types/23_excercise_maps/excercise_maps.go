package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount 単語数集計処理
func WordCount(s string) map[string]int {
	// mapはゼロ値がnilなので、明示的に初期化
	wordCount := map[string]int{}
	// strings.Split(s, sep)で文字列を配列へ変換。
	splitArray := strings.Split(s, " ")
	for _, word := range splitArray {
		value := wordCount[word]
		wordCount[word] = value + 1

	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
