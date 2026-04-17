package mutexes

import (
	"fmt"
	"sync"
)

func Increment(wg *sync.WaitGroup, lock *sync.Mutex, counter *int) {
	defer wg.Done()
	lock.Lock()
	*counter++
	value := *counter
	lock.Unlock()
	fmt.Printf("Incrementing: %d\n", value)
}

func Decrement(wg *sync.WaitGroup, lock *sync.Mutex, counter *int) {
	defer wg.Done()
	lock.Lock()
	*counter--
	value := *counter
	lock.Unlock()
	fmt.Printf("Derementing: %d\n", value)
}
