package main

import (
	"fmt"
	"time"
)

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
var logCh1 = make(chan logEntry, 50)
var doneCh = make(chan struct{})
// ** signal only channel **
// struct with no field requires zero memory allocation
// useful to identifying when a channel is closed


func main() {
	go logger() // monitors log channel for entries
	defer func(){
		close(logCh)
	// can use defer to close channel after main func exits
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)

	go logger1()
	logCh1 <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh1 <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger(){
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

func logger1(){
	for {
		select {
// select block breaks loops when empy struct is sent
// to doneCh on line 42
		case entry := <- logCh1:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <- doneCh:
			break
		}
	}	
}

/*
Summary
- Create channel with make()
- Must be strongly typed
- Send and receive with <-
- Can have multiple data senders and receivers
- Buffered channel with second int argument; buffered channels:
	- block sender side till receiver is available
	- block receiver side will message is available
	- can decouple sender/receiver
	- be used when send/receive have assymmetric loading
- For range loops
	- the first parameter is the value not index
- Select statements are used to monitor the state of a channels
	- can use default case if channel blocking is not required
*/