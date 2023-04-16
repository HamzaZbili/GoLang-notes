package main

import (
	"fmt"
)

// application starts in main function
func main(){
	// main function does not take parameters
	// upper and lower case determines visibility
	sayMessage("hello there")
	name := "Hamza"
	surname := "Zbili"
	sayGreeting("hello", name)
	fmt.Println(name) // still "Hamza"

	printSurname(&surname) // can pass pointer as param

	sum("The sum is:",1,2,3,4,5,6,7,8,9)

	s := sum2(2,3465,4367,374,2,32,32)
	fmt.Println(s)

	s1 := sum3(32135513,1513553,121356,323)
	fmt.Println(s1) // receives address
	fmt.Println(*s1) // receives address and dereferences

	s2, err1 := div(8.342432, 2.5253252335)
	if err1 != nil {
		fmt.Println("s2", err1)
	}
	fmt.Println("s2", s2)
	s3, err2 := div(5.0, 0.0)
	if err2 != nil {
		fmt.Println("s3",err2)
	}
	fmt.Println("s3",s3)
}

// functions are hoisted
func sayMessage(msg string){ // params are strong typed
	fmt.Println(msg)
}

func sayGreeting(greeting, name string){
// Go assumes each var in comma delimited list is same type
	fmt.Println(greeting, name)
// changing variables here won't change variable in main func
	name = "Ted"
	fmt.Println(greeting, name) // but will in sayGreeting func
}

func printSurname(surname *string){ // receives pointer
	fmt.Println(surname) // prints location -> 0x1400010a240
	fmt.Println(*surname) // dereferences location -> Zbili
	// this is useful for perfomance benifits
	// cannot do this for maps and slices
}

func sum(msg string, values ...int){ // variadic param -> wraps values in slice
	// variadic params have to be at the end and there can only be one
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func sum2(values ...int) int { // return type declared here
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	return result // if return type assigned, must return
}

func sum3(values ...int) *int { // functions can return pointers
	fmt.Println(values)
	// result variable is declared on execution stack of this function
	result := 0
	// when function exits, that memory is freed up
	for _,v := range values {
		result += v
	}
	// Go recognises return value is generated on local stack
	// and promotes variable to heap memory (shared)
	return &result // returns address of result
}

func sum4(values ...int) (result int) {
	 // can also declare result variable here
	fmt.Println(values)
	for _,v := range values {
		result += v
	}
	return // results result variable as before

	// this isn't common syntax
}

func div(a, b float64) (float64, error) { 
	// can return multiple values (eg error)
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	} // conditional handles error
	return a / b, nil
}


