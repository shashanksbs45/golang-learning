//In Go, Goexit is a function provided by the "runtime" package, and it is used to terminate the goroutine in which it is called.
//When Goexit is called, the current goroutine is terminated, but other goroutines continue to execute. Below is an example demonstrating the use of Goexit:

package main

import (
	"fmt"
	"runtime"
	"time"
)

func GoExit() {
	// Create a goroutine
	go func() {
		defer fmt.Println("Deferred function in the goroutine")
		fmt.Println("Goroutine will exit")
		runtime.Goexit() // Terminate the current goroutine
	}()

	// Allow time for the goroutine to execute
	time.Sleep(1 * time.Second)

	fmt.Println("Main function will continue")
}
