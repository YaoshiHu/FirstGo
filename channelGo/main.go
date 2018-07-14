package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "message 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "message 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("received: ", msg1)
			case msg2 := <-c2:
				fmt.Println("received: ", msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
