package main

// imports here
import (
	"fmt"
	"strconv"
)

// var block used to declare mutliple variables
var (
	// lowercase  variables are scoped to package
	// upper are global
	// variables declared in block are only retrievable inside block
	name = "hamza"
	// the length of the variable name should reflect lifespan
	age = 30
)
// also declarable like this - useful for declaring before assigning
var surname string = "zbili"




func main(){
	// below declares and initialises (only in block)
	foo := "this declares and initialises"
	// println to print in console
println(name, age)
	// inside block - can declare like this 
height := 180

println(height, surname)
// below to check variable type and value
fmt.Printf("%v, %T", name, age)
var j string
// use strconv to convert int to string
// import above required
j = strconv.Itoa(age)
println("%v, %T", j, j)
println(foo)
}
  