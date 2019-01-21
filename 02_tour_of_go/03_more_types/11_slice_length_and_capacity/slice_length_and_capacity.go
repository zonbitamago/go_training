package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// lengthが減る
	s = s[:0]
	printSlice(s)

	// lengthが増える
	s = s[:4]
	printSlice(s)

	// capacityが減る
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
