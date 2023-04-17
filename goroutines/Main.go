package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main () {
	// most other languages use OS threading

	// a goroutine is a thread execution that can run
	// concurrently with the main function


	go sayHello() // goroutine initialised like this
	// this will spawn a new goroutine that executes sayHello
	time.Sleep(100 * time.Millisecond)
	// the main function will exit before sayHello has the chance to return

	
	var msg = "hi there"
	go func(){ // anon func has access to var in outer scope
		fmt.Println(msg)
	}()
	msg = "goodbye" // race condition
	// var is reassigned before the goroutine has a chance to print
	time.Sleep(100 * time.Millisecond)

	var msg1 = "hello there"
	go func(msg1 string){
		fmt.Println(msg1)
	}(msg1) // passing the variable like this avoids raise condition
	msg1 = "goodbye"
	time.Sleep(100 * time.Millisecond)

	// wait group -> alternative to time.Sleep
	msg2 := "first"
	wg.Add(1) // synronises wait group
	go func(msg2 string) {
		fmt.Println(msg2)
		wg.Done() // decrements wait group
		// this method doesn't rely on real time
	}(msg2)
	msg2 = "second"
	wg.Wait()  // exits goroutine
	
	

}

func sayHello(){
	fmt.Println("hello")
}