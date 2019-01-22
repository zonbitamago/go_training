package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// スライスの末尾に0を追加
	s = append(s, 0)
	printSlice(s)

	// スライスの末尾に1を追加
	s = append(s, 1)
	printSlice(s)

	// スライスの末尾に2,3,4をまとめて追加
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
