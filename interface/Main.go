package main

import (
	"bytes"
	"fmt"
)

func main(){
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello go"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	// Write method expects byte slice
	wc.Write([]byte("oewifownfgiwnfjnwqefijnewijnfijne"))
	wc.Close()

	// empty interface has no methods
	var myObj interface{} = NewBufferedWriterCloser()
	// anything can be cast to an empty interface
	// can be used to apply an interface to different types
	if bwc, ok := myObj.(WriterCloser); ok {
		bwc.Write([]byte("emptyinterface"))
		bwc.Close()
	}
	// may need to do a type conversion or use reflect package
	// to identify the types that are being utilised
}

// an interface is like contract between different parts
// of a program specifying a common set of behaviors that can
// be expected from any type that implements the interface

// naming convention -> end in er
type Writer interface { // declared like struct
	// Write method from io package
	// accepts slice of bytes and returns int and error
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write (data []byte) (int,error) {
// here ConsoleWriter implements the Writer interface
// therefore thhe ConsoleWriter struct is an instance
// of the Writer interface
	n, err := fmt.Println(string(data))
	return n, err
	
}

type Incrementer interface {
	Increment() int // Increment method has to return integer
 }

 type IntCounter int // define type alias

 func (ic *IntCounter) Increment() int { // IntCounter receives int
	*ic++ // implementation increments type itself
	return int(*ic) // 
 }

 type Closer interface {
	Close() error
 }

 // creates interface composed of other interfaces
 type WriterCloser interface {
	Writer
	Closer
 }

 // example -> 
 type BufferWriterCloser struct {
	buffer *bytes.Buffer
 }

 // implement Write method
 func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) // assigns input to n and err
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
 }

  // implement Close method
 func (bwc *BufferWriterCloser) Close() (error) {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
 }

 // constructor function returns pointer buffered writer closer
 func NewBufferedWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
 }

