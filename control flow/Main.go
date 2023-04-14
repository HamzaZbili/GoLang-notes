package main

import (
	"fmt"
	"math"
)

func main() {
	// structure of if statement
	if true { // literal true
		fmt.Println("this test is true")
	}
	countryPopulations := map[string]int{
		"Italy": 3446334,
		"France": 455424,
		"UK": 6235232,
		"Germany": 823432,
		"Spain": 9234423,
	}
	if pop, ok := countryPopulations["Italy"]; ok {
		// initialiser syntax
		// the above is a call to a map
		// as ok is true the statement executes
		fmt.Println(pop)
		fmt.Println(ok)
	}
	// comparison operators
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("too low")
	}
	if guess > number {
		fmt.Println("too high")
	}
	if guess == number {
		fmt.Println("that's correct")
	}
	fmt.Println(number<=guess, number>=guess, number!=guess)
	// and operator -> both conditions must be met
	if guess<number && guess>40{
		fmt.Println("almost there")
	}
	// or operator
	if guess>1 || guess<100 {
		// or operator -> will run when one is met
		// short circuits when first true
		fmt.Println("guess has to be between 1 and 100")
	} else if guess != number{
		fmt.Println("in the right area")
	} else {
		fmt.Println("that's it!")
	}
	num := 0.123
	// when working with floating point numbers
	// go returns an approx result which may not be reliable
	if num == math.Pow(math.Sqrt(num), 2) {
		fmt.Println("these are the same")
	} else {
		fmt.Println("these are different")
	}

	// switch statement
	switch 2 {
		// first true case will execute
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("neither one or two")
	}

	switch i := 2 + 3; i {
		// first true case will execute
	case 1, 3, 5, 7, 9:
		fmt.Println("odd smaller than ten")
		// break is implied implicitally
	case 2, 4, 6, 8:
		fmt.Println("even smaller than 10")
	default:
		fmt.Println("another number")
	}

	// type switching 
	var i interface{} = 1
	// type variable is assigned to interface using below syntax
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}
