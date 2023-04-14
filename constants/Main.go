package main

import (
	"fmt"
)

const a int16 = 13

func main (){
	// declared with lower case to be used in file as before
	// value ust be calcuble at compile time
	const speedOfLight int = 78937
	fmt.Printf("%v, %T\n",speedOfLight, speedOfLight);
	// arrays cannot be declared as const
	
	const a int = 18;
	// immutable but can be shadowed
	fmt.Printf("%v, %T\n",a, a)

	const b  = 3
	var c int16 = 4
	// var and const can be added together (makes a var)
	// if one int is unnasigned the above works
	fmt.Printf("%v, %T\n", b+c, b+c)

	const (
		_ = iota
		// _ right only variable in Go
		i = iota
		j = iota
		k = iota
	)
	// this will print 0, 1, 2 -- before addition of _
	// iota value changes as the constants are being evaluated
	// this is specific to the const block
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	// cannot apply ^ to iota -> but can use bitshift

	const  (
		_ = iota // ignores first value
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fileSize := 5000000000.
	fmt.Printf("%.2fGB\n", fileSize / GB)



}