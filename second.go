package main

import "fmt"

func main() {
	f()
	fmt.Printf("End of Program")
}

func f() {
	defer func() {
		fmt.Println("Defer start")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Defer end")
	}()
	for {
		fmt.Println("Start of f()")
		a := []string {"a", "b"}
		fmt.Println(a[3])
		panic("Fake panic")
		fmt.Println("End of f()")
	}
}
