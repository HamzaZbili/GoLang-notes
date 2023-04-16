package main

import (
	"fmt"
)

func main () {
	a := 42
	b := a // data is coppied and assigned to b
	fmt.Println(a, b) // 42 42
	a = 27 // changes a but not b
	fmt.Println(a, b) // 27 42

	var c int = 31
	// *int -> declares point to an int
	var d *int = &c //& -> "address of" operator
	fmt.Println(c, d)
	// d is not holding value 31
	// it holds the memory location of c (0x140000ac010)
	fmt.Println(&c, d) // 0x140000ac010 0x140000ac010

	// * dereferences the pointer
	fmt.Println(*d) //prints value (31)

	arr := [3]int32{12,23,34}
	poi0 := &arr[0] // points to first value in arr
	poi1 := &arr[1] // points to second value in arr
	poi2 := &arr[2] // points to second value in arr
	fmt.Printf("%v, %p, %p, %p\n", arr ,poi0, poi1, poi2)
	//[12 23 34], 0x140000ac030, 0x140000ac034, 0x140000ac038
	// location of each value is 4 bytes apart in Go array
	// Go does not support pointer arithmetics
	// "unsafe" package allows you to overwride this **avoid**

	var ss myStruct // declares struct
	ss = myStruct{foo: 42} // assign value
	fmt.Println(ss) // {42} 

	var ms *myStruct //* declares pointer to myStruct object
	ms = &myStruct{foo: 42} // & (address of) on object initialiser
	// ms is holding address of object with value 42
	fmt.Println(ms) // &{42} 

	var oo *myStruct
	// unintialised pointer var has value <nil>
	fmt.Println(oo) // <nil>
	// can also use new() to initialise var pointer 
	oo = new(myStruct) // new() can only initialise empty object
	fmt.Println(oo) // &{0} 
	
	var pp *myStruct
	pp = new(myStruct)
	// to get to unintialised point (<nil>)
	pp.foo = 42 // dereferences pp var and assigns value
// (*pp).foo --> compiler knows to target object over pointer
	fmt.Println(pp.foo) // 42

	sli := []int{1,2,3}
	slipoi := sli
	// slices are projections of underlying arrays
	// - do not hold data
	// - holds pointer to first element of underlying array
	fmt.Println(sli,slipoi) // [1 42 3]
	sli[1] = 42
	fmt.Println(sli,slipoi) // [1 42 3]

	// maps have the same behaviour as slices
	ma := map[string]string {"foo": "bar", "baz": "buz"}
	mb := ma
	fmt.Println(ma, mb) // map[baz:buz foo:bar] map[baz:buz foo:bar]
	// both maps will be modified
	ma["foo"] = "quz"
	fmt.Println(ma, mb) // map[baz:buz foo:quz] map[baz:buz foo:quz]

	// ***** summary ******
	// - * -> declare type to pointer to underlying data
	// - & -> addresof operator to get address of variable
	// - * in front of pointer var -> dereferences pointer
	// - complex types (eg structs) are auto dereferenced
	// 		- (*pp).foo -> pp.foo
	// objects -
	// - point := &object -> if address if value ex///
	// - new(myStruct) - keyword creates object wi/thout initialising fields
	// all assignment operations in Go are copy operations
	// slices and maps copy pointers not underlying data
}

type myStruct struct{
	foo int
}
