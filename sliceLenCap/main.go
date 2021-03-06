package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = append(s, 2, 4)
	printSlice(s)

	s = s[:cap(s)]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

