package main

import (
	"fmt"
	"runtime" // provides low-level access to
	// Go runtime system which is responsible for:
	// - managing goroutines
	// - memory allocation
	// - garbage collection
	// - other system-levelfunctions
)

func main () {
	/* GOMAXPROCS is an environment variable and
	a function that specifies the maximum number of
	operating system threads that can execute
	user-level Go code simultaneously.*/

	// GOMAXPROCS returns the number threads that were previously set
	// -1 required to receive number of available threads
	fmt.Printf("Threads: %v/n", runtime.GOMAXPROCS(-1))

	runtime.GOMAXPROCS(100) // sets available threads
	// increasing threads:
	// improves the parallelism of a program, especially
	// if it is CPU-bound and has many goroutines that
	// can run concurrently
	fmt.Printf("Threads: %v/n", runtime.GOMAXPROCS(-1))
	// setting GOMAXPROCS too high (eg 100) will cause
	// cause overhead issues
}

/* 
**** Best Practices ****
- Avoid goroutines in libraries
- Know when a goroutine will end on creation
- Check for race condition on compile time
	$ go run -race Main.go -> finds data race(s)

*/