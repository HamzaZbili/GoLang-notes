package main

import (
	"fmt"
)


func main(){
	func() { 
	// anonymous function is executed on declaration
		fmt.Println("hello go")
	// may be useful when assigning variables for single use
	}()

	for i := 0; i < 6; i++ {
		func(i int) { // provide variable here
			fmt.Println(i) // can access i from loop container
		}(i) // pass variable here -> important for async funcs
	}

	f := func() {
	fmt.Println("testing")
	} // assign anonymous function to variable
	f()

	var fu func () = func () {
		fmt.Println("this is the same as above")
	}
	fu()

	var divide func (float64, float64) (float64, error)
	divide = func(a,b float64) (float64, error) {
		// same as before, anon func with two float64
		if b == 0.0 {
			// checks not trying to divide by 0.0 and return error
			return 0.0, fmt.Errorf("cannot divde by zero")
		} else {
			return a / b, nil // returns result and nil
		}
	}
	firstDiv, err1 := divide(812821.23, 213123.1423)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(firstDiv)
	secondDiv, err2 := divide(9.2, 0.0)
	if err2 != nil {
		fmt.Println(err2)
		// return (commented so main does not exit)
	}
	fmt.Println(secondDiv)
	// when functions are assigned to variables
	// they are no longer hoisted

	// declares greeter object from struct below
	g := greeter {
		greeting: "hello",
		name: "go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name string
}
// **** methods *****
func (g greeter) greet() { // get access to struct and g object
	// "greeter" is receiver
	// method is function that executes in known context
	fmt.Println(g.greeting, g.name)
	// methods work with pointer -> original data not modified
	// if method reassigns value eg. g.greeting = "something else"

	// by using * and & -> can gain access to underlying value
}

// methods can be used on any type

/* **** Summary ****
- functions can be exported with upper case
- params of same type can be comma delimited with
- if pointers passed
	- functions can modify value with * / &
	- always the case with slices and maps
- variadic params sent sent with ...
	- must be last param
	- will receive slice
- can return multiple values (eg error)
	- if assigning multiple return values to variable, add second variable
- addresses of local variables can be returned
	 - Go will automatically promote returned addresses to local (heap) from shared (stack)
- anonymous functions have to be declared before they are assigned to a variable
- methods execute in context of a type (often struct)
	- value receiver gets copy of type
	- pointer receiver gets pointer to type (add * before type)
*/