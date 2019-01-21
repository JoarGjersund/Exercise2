// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
			case buff := <-add_number:
				i+=buff;
				break
			case temp := <-control: // control signals will lead us here:
				switch temp {
					case GetNumber:
						Println("GetNumber, ok.")
						number <- i;
					case Exit:
						Println("Exit, ok.")
						number <- 0;
						return
				}			
		}
	}
	
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}
	finished<-true
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- -1
	}
	finished<-true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	number := make(chan int) // channel for the final number
	add_number := make(chan int,0) // channel for the number to be added.
	finished := make(chan bool) // channel for workers to signal when finished.
	control := make(chan int) // channel to control the number server
	
	
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	// answ: the buffer size of add_number channel must be 0 (default), if it is more the worker routines are not properly synchronized anymore.

	// Spawn the required goroutines
	go decrementing(add_number,finished)
	go incrementing(add_number,finished)
	go number_server(add_number,control,number)

	<-finished // block until finished from the first worker
	<-finished // block until finished from the last worker.
	control<-GetNumber 
	
	Println("The magic number is:", <- number)

	control<-Exit
	<-number // wait until program is exited.
}
