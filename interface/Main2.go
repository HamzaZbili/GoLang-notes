package main

import (
	"fmt"
	"math"
)

type shape interface {
	// each struct that has the method below
	// qualifies as a type shape
	area() float64
	// pow() -> adding a method would
	// disqualify circle and rect
}

type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

type triangle struct {
	width float64
	height float64
}

func (r rect) area() float64 {
	// defines method area() for rect
	return r.width * r.height
}

func (c circle) area() float64 {
	// defgines method area() for circle
	return math.Pi * c.radius * c.radius
}

func (t *triangle) area() float64 {
	return .5 * t.width * t.height
}
 


func getArea ( s shape) float64 {
	// function wraps interface methods
	return s.area()
}

func main() {
	r1 := rect{5,6} // define rect
	c1 := circle{4.5} // defines circle
	t1 := triangle{4,10}
	// create array of shapes
	shapes := []shape{c1, r1, &t1}
	// as triangle receives pointer -> requires addressOf

	// can only access the methods defined in interface
	for _, shape := range shapes {
		// loops through array
		// prints area of shape
		fmt.Println(shape.area())
	}

	fmt.Println(shapes[1].area()) // returns same as above

	/*	 **** pointer / receiver ****
	 when implementing an interface
	 -if implemented with value type ->
	the methods have to all have value receivers
	-if  implemented with pointer type -> 
	the methods can have both value and pointer receivers */

	

}