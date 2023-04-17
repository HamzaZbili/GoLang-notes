package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <- chan int) {
		// ensures channel is receive only
		i := <- ch 
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		// ensures channel is send only
		i := 42
		ch <- i
		wg.Done()
	}(ch)
	wg.Wait()
}