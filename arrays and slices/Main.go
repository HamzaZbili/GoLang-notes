package main

import (
	"fmt"
)

func main () {
	// must declare size and types
	// each element in array is contiguous in memory
	grades := [...]int{87,45,65}
	// by adding ... the array will be correct size
	// can replace with int if values are not added
	fmt.Printf("Grades: %v\n",grades)

	// can also be declared like this
	// cannot be const 
	var names [3]string
	names[0] = "Shmed"
	names[2] = "Sasha"
	names[1] = "Sarah"
	fmt.Printf("Names: %v\n",names)
	fmt.Println(len(names))

	// can create array of arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1,0,0}
	identityMatrix[1] = [3]int{0,1,0}
	identityMatrix[2] = [3]int{0,0,1}
	fmt.Println(identityMatrix)

	a := [...]int{1,2,3}
	// b creates a new array
	b := a
	// if modified, can still access a and b
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// pointers are used to point to original data
	// & -> addressOf operator
	c := &a
	// this will change both a and c
	c[0] = 9
	fmt.Println(a) // [0 5 3] === array itself
	fmt.Println(c) // &[0 5 3] ==== pointing to a

	// slices are declared like this -> illiminating ...
	d := []int{1,2,3}
	fmt.Println(d)
	// as length is not declared in slice
	// we have access to cap()
	fmt.Println(cap(d))

	arr := []int{1,2,3,4,5,6,7,8,9}
	fullslice := arr[:]
	fromslice := arr[3:]
	uptoslice := arr[:6]
	betweenslice := arr [3:6]

	// all the above point to the same underlining arr
	// therefore the below modifies all the above
	arr[5] = 42
	
	fmt.Println(arr)
	fmt.Println(fullslice)
	fmt.Println(fromslice)
	fmt.Println(uptoslice)
	fmt.Println(betweenslice)

	// make function takes 2 or 3 arguments
	m := make([]int,3, 100) // array or ints -> 3 elements len
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	// cap function returns capacity of underlying array
	fmt.Printf("Capacity: %v\n", cap(m))

	n := []int{}
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))
	// unlike arrays, slices do not have to have a fixed size
	// ie we can add and remove elements
	n = append(n, 1)
	// above adds element to index
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	// concat slices like this
	n = append(n, []int{2,3,4,5,6,7,8,9}...)
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))


	var k []int = n[1:] // removes first
	fmt.Println(k)
	k = n[:len(n) -1] // removes last
	fmt.Println(k)
	fmt.Println(n)
	k = append(n[:1], n[8:]...) // removes from inside
	fmt.Println("n is", n)
	// careful with this and it will modify original
	fmt.Println("k is", k)

	// slices are more flexible and commonly used in
	// Go because of their dynamic size and mutability
	// but arrays are useful in certain cases where a
	// fixed size is needed or when passing the array
	// to a function without the possibility of
	// modification is needed
}