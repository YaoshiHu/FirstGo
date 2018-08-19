package main

import "fmt"

func main() {
	const (
			Red = iota
			Green
			Blue
		)
	fmt.Println("Red:", Red, " Green:", Green, " Blue:", Blue)
}