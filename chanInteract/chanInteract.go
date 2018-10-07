package main

import "fmt"

func main() {
	numChan := make(chan int, 1)
	doneChan := make(chan struct{})

	go func() {
		fmt.Println(<-numChan+1)
		doneChan <- struct{}{}
	}()

	numChan<-8
	<-doneChan
}

