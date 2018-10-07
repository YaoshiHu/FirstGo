package main

import (
	"fmt"
	"sync"
	"net/http"
)

func printStatusCode(pid int, url string) {
	res, _ := http.Get(url)
	fmt.Printf("%d, %d\n", pid, res.StatusCode)
}

func main() {
	var wg sync.WaitGroup
	urls := [...]string{
                "https://google.com",
                "https://facebook.com",
                "https://dev.to",
                // ... thousands more
        }
        for i, url := range urls {
                wg.Add(1)
                go func(pid int, url string) {
                        printStatusCode(pid, url)
                        wg.Done()
                }(i, url)
        }
        wg.Wait()
}
