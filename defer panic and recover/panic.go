package main

import (
	"fmt"
	"net/http"
)

func main () {
		// ***** panic *******

	// i, k := 1, 0
	// ans := i / k
	// fmt.Println(ans)
	// will generate panic
	// panic: runtime error: integer divide by zero

	fmt.Println("test1")
	// can also call a panic and custom string
	// ======= panic("something bad happened")
	// this will not print "test2"
	fmt.Println("test2")

	// http.HandleFunc -> webrequest
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// response write gives access to webrequest
		w.Write([]byte("Hello!"))
	})
	// error variable checks if port is blocked
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		// conditional throws error
		panic(er.Error())
	}

	// panics do not have to be fatal
	// they are if panic all the way up to
	// the go runtime
}