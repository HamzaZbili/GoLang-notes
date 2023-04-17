package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ { 
// creates 10 sets of gorountines (senders/receivers)
		wg.Add(2)
		go func() {
			j := <- ch
			fmt.Println(j)
			wg.Done()
		}()
		go func() {
// only one message can be in an
// unbuffered channel at a time
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}