package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var wg1 = sync.WaitGroup{}
var wg2 = sync.WaitGroup{}


func main() {
	ch := make(chan int, 2)
	// buffer channel by adding second int argument 
	wg.Add(2)
	go func(ch <- chan int) {
		i := <- ch 
		fmt.Println(i)
		j := <- ch
		fmt.Println(j)
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		i := 42
		j := 30
		ch <- i
		ch <- j
		wg.Done()
	}(ch)
	wg.Wait()

	ch1 := make(chan int, 2)
	// buffer channel by adding second int argument 
	wg1.Add(2)
	go func(ch1 <- chan int) {
		for i := range ch1 {
			fmt.Println(i)
		}
		wg1.Done()
	}(ch1)
	go func(ch1 chan <- int) {
		ch1 <- 99
		ch1 <- 1
		close(ch1)
	// closing the channel informs receiver
	// that it's done sending
		wg1.Done()
	}(ch1)
	wg1.Wait()

	ch2 := make(chan int, 2)
	wg2.Add(2)
	go func(ch2 <- chan int) {
		for {
	// conditional checks if channel is open
	// and does not require for range loop
			if i, ok := <- ch2; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg2.Done()
	}(ch2)
	go func(ch2 chan <- int) {
		ch2 <- 734
		ch2 <- 923
		close(ch2)
		wg2.Done()
	}(ch2)
	wg2.Wait()
}