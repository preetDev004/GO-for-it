package main

import (
	"fmt"
	"sync"
)

// For reading the data in multiple go routines we can use (sync.RWMutex) RLock() and RUnlock()  so the data will be available for 
// reading but when you write you must use (sync.Mutex) Lock() and Unlock().

func main() {
	// Create a new shared variable
	var count int

	// Create a new sync.Mutex
	var mutex sync.Mutex

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Add 10 goroutines to the WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		// Start a new goroutine
		go func() {
			// Decrement the WaitGroup counter when the goroutine finishes
			defer wg.Done()

			// Lock the mutex before accessing the shared variable
			mutex.Lock()
			defer mutex.Unlock()

			// Increment the shared variable
			count++

			// Print the current value of the shared variable
			fmt.Printf("Count: %d\n", count)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}