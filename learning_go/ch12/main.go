package main

import (
	c "ch12/channels"
	r "ch12/routine"
	"fmt"
	"sync"
	"time"
)

func main() {

	r.Routine()
	oneToTen := r.IterationParams{Start: 1, Stop: 10, Step: 1}
	oneToHundred := r.IterationParams{Start: 1, Stop: 100, Step: 1}

	{ // using goroutines with time.Sleep
		var sumResult1, sumResult2 int
		go func() {
			sumResult1 = r.Summation(oneToTen)
		}()
		go func() {
			sumResult2 = r.Summation(oneToHundred)
		}()
		time.Sleep(time.Second * 1)
		fmt.Println(sumResult1, sumResult2)
	}
	{ // using goroutines with sync.WaitGroup
		var sumResult1, sumResult2 int
		var wg sync.WaitGroup
		wg.Add(2)
		go r.SummationWithWaitGroup(oneToTen, &wg, &sumResult1)
		go r.SummationWithWaitGroup(oneToHundred, &wg, &sumResult2)
		wg.Wait()
		fmt.Println(sumResult1, sumResult2)
	}
	{ // using goroutines with sync/atomic
		var sumResult int32
		var wg sync.WaitGroup
		wg.Add(4)
		go r.SummationWithAtomicOperations(
			struct{ Start, Stop, Step int }{1, 25, 1}, &wg, &sumResult)
		go r.SummationWithAtomicOperations(
			struct{ Start, Stop, Step int }{26, 50, 1}, &wg, &sumResult)
		go r.SummationWithAtomicOperations(
			struct{ Start, Stop, Step int }{51, 75, 1}, &wg, &sumResult)
		go r.SummationWithAtomicOperations(
			struct{ Start, Stop, Step int }{76, 100, 1}, &wg, &sumResult)
		wg.Wait()
		fmt.Println(sumResult)
	}
	{ // using goroutines with string result
		var sumResult string
		var wg sync.WaitGroup
		wg.Add(4)
		go r.SummationWithString(1, 25, &wg, &sumResult)
		go r.SummationWithString(26, 50, &wg, &sumResult)
		go r.SummationWithString(51, 75, &wg, &sumResult)
		go r.SummationWithString(76, 100, &wg, &sumResult)
		wg.Wait()
		fmt.Println(sumResult)
	}

	{
		// using goroutines with channels
		ch := make(chan string)
		go c.Greet(ch)
		ch <- "Hello, John Travotal"
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}
}
