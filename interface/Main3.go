package main

import (
	"fmt"
)

func main () {
	// empty interface used to determine type
	var i interface{} = 0
	switch i.(type) {
	case int: 
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is an string")
	default:
		fmt.Println("i is not known")
	}

/* BEST PRACTICES
	- use many small interfaces
	- io.Writer, io.Reader, interface{} -> from Go library
	- don't export intefaces for types that will be consumed
	- export interfaces for types that will be used by package
	- design functions and methods to receive interfaces when possible
*/
}