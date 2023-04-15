package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main () {
	// ****** defer ******
	fmt.Println("first")
	// defer keyword is called last
	defer fmt.Println("second")
	fmt.Println("third") // will print first third second

	// resources are closed in the opposite order
	// after the main functions but before the functions returns
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3) // will print 3 2 1

	res, err := http.Get("http://www.gooogle.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer can be used to close a resource
	// next to the request as below
	defer res.Body.Close()
	// read response and assign to variable
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// checking read operation
		log.Fatal(err)
	}
	// print response below then close resourse above
	log.Printf("%s", robots)
	//  ******  defer in loops *****
	// if making a lot of requests within a loop
	// because the defer statement doesn't execute
	// until all loops are closed
	// -- assign the closing of the resource to a function
	// and call that function where you want the resource
	// to close

	a := "start"
	defer fmt.Println(a) // this will print a
	// defer takes the value of variable at line of declaration
	a = "end"
}