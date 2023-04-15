package main

import (
	"fmt"
	"log"
)

func main () {
	fmt.Println("start")
	
	defer fmt.Println("end")
	// panics happen after defer resources
	// so "end will print" despite line below
	// ** panic("something bad happened") ** 
	
	// order of execution is:
	// - main function
	// - deferred statements
	// - panics
	// - return value

	fmt.Println("first")
	// anonymous function can only be called once
	defer func() {
		// defer statement takes a function call
		// not a function itself
		if err := recover(); err != nil {
		// recover function returns nil if app panics
			log.Println("Error:", err)
		}
	}()
	// panic("something bad happened") -> second won't print
	fmt.Println("second")

	fmt.Println("hello")
	panicker() // main function continues
	// despite panic thrown inside panicker() function
	fmt.Println("there")

	// summary
	// panics are used when an application gets into
	// state that it can't recover from

	// recover() is used to recover from panis
	// - only useful in deferred functions
	// - the currect function will not continue;
	// - but rest higher functions will after recovered
}


func panicker() {
	fmt.Println("about to panic")
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				// panic(err)
		// adding panic(err) will throw panic in main func
			}
		}()
		// thorws panic
	panic("something bad happened")
	// fmt.Println("panic over") // does not print
}