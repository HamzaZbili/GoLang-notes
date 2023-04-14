package main

import (
	"fmt"
)

func main() {
	var boo bool = true
	fmt.Println(boo)


	// below logical statement
	a := 1 == 1
	b := 1 == 2
	// print true boolean
	fmt.Printf("%v, %T", a, a)
	// prints false boolean
	fmt.Printf("%v, %T", b, b)

	//// numeric types

	// default integer types
	m := 42
	// prints 42 int
	fmt.Printf("%v, %T", m, m)
	// general int - must specify if using large numbers
	// other int types are
	// signed integers are used for negative and positive number
	// int8 int16 int32 int64

	var n uint16 = 54
	// unsinged integers start from 0 and up
	fmt.Printf("%v, %T", n,n)

	int1 := 10 // 1010
	int2 := 3  // 0011

	fmt.Println(int1 / int2)
	// prints 3 as type cannot change
	fmt.Println(int1 % int2);
	// use above to find remainder
	
	// cannot preform operations on two different integer types
	// fmt.Println(int1 + n)

	// bit operators
	// 10 = 1010
	//  3 = 0011
	fmt.Println(int1 & int2)  // and returns 0010 (2)  // both must be set
	fmt.Println(int1 | int2)  // or  returns 1011 (11) // either can be set
	fmt.Println(int1 ^ int2)  // xor returns 1001 (9)  // either can be set, but not both
	fmt.Println(int1 &^ int2);// not returns 0100 (4)  // both most not be set

	// bit shift
	fmt.Println(int1 << 1) // each bit shifts to left  - 10100 (20)
	fmt.Println(int1 >> 1) // each bit shifts to right - 00101 (5)

	// floating point numbers
	float :=  3.14
	float = 13.7e72 
	float = 2.1E14
	// all of the above are valid floating number declarations
	fmt.Printf("%v, %T", float, float)
	// bit operators, bit shift operators and remainder operators
	// are not available when using floating point nunmbers

	var imagine complex128 = 1 + 1i
	// Go recognises i and is usable in operations
	fmt.Printf("%v, %T\n", real(imagine), real(imagine))
	fmt.Printf("%v, %T\n", imag(imagine), imag(imagine))

	//// text types

	s := "this is a string"
	// strings are like arrays in Go however they are immutable
	// therefore 2[1] = "j" will through an error

	// strings only support UTF-8
	fmt.Printf("%v, %T", s ,s)
	// below is a convertion to a slice of bytes
	stringbite := []byte(s)
	// this is useful for sending data/files
	fmt.Printf("%v, %T", stringbite ,stringbite)
	// the above can then be converted back

	// rune - UTF-32
	// declared with single quotes
	var r rune = 'a'
	// runes are a true type alias
	// the value of r IS 97 and not 'a'
	fmt.Printf("%v, %T", r ,r)
}