package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

var wg1 = sync.WaitGroup{}
var counter1 = 0
var m = sync.RWMutex{}

func main() {
	for i:= 0; i < 10; i++ {
		// spawns 20 goroutines
		wg.Add(2)
		go sayHello()
		go increment()
		// goroutines run at the same time and won't
		// return syncronous result
	}
	wg.Wait()


	// Mutex is a lock that the application will honour
	for i:= 0; i < 10; i++ {
		// locking mutex ensures synchronisation 
		wg1.Add(2)
		// locks mutex outside context of gorountine
		m.RLock()
		// . RLock() used to aquire a shared lock
		go sayHelloM()
		// .Lock() used to aquire an exclusive lock
		m.Lock()
		go incrementM()
	}
	wg1.Wait()


}

func sayHello(){
	fmt.Println("Counter:", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func sayHelloM(){
	fmt.Println("Counter mutex:", counter)
	// unlocks shared mutex lock
	m.RUnlock() 
	wg1.Done()
}

func incrementM() {
	counter++
	// unlocks exclusive mutex lock
	m.Unlock()
	wg1.Done()
}