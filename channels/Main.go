package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	ch := make(chan int) // declared with make(chan)
	// type specifies type that can be sent/received by channel
	waitGroup.Add(2)
	go func() {
		// value received using := <-
		i := <- ch 
		fmt.Println(i)
		waitGroup.Done()
	}()
	go func() {
		i := 42
		ch <- i // values sent to cannel using <-
		i = 40 
	// sends copy so reassigning value as above
	// ensures previous value is received
		waitGroup.Done()
	}()
	waitGroup.Wait()
}