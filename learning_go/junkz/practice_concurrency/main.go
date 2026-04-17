package main

import (
	m "practice_concurrency/mutexes"
	"sync"
)

func main() {

	//a.RunNonAnons()
	//a.RunAnons()

	// using Mutex
	{
		var wg sync.WaitGroup
		var lock sync.Mutex
		var counter int
		for i := 0; i < 5; i++ {
			wg.Add(2)
			go m.Increment(&wg, &lock, &counter)
			go m.Decrement(&wg, &lock, &counter)
		}
		wg.Wait()
	}
}
