package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared counter that will be accessed by multiple goroutines
var counter int

// Mutex to ensure that only one goroutine can access the counter at a time
var mutex sync.Mutex

// Function to increment the counter
// Accepts a WaitGroup pointer and a unique ID for each goroutine
func increment(wg *sync.WaitGroup, id int) {
	// Defer marks the WaitGroup as done when this function exits
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// Lock the mutex to gain exclusive access to the counter
		mutex.Lock()

		// Critical section: increment the shared counter
		counter++
		fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)

		// Unlock the mutex to allow other goroutines to access the counter
		mutex.Unlock()

		// Sleep for a short duration to simulate work and make the output more understandable
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines we want to spawn
	numGoroutines := 3

	// Launch multiple goroutines
	for i := 1; i <= numGoroutines; i++ {
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)

		// Launch a goroutine to run the increment function
		go increment(&wg, i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final value of the counter after all goroutines are done
	fmt.Printf("Final counter value: %d\n", counter)
}
