package channels

import (
	"fmt"
	"time"
)

func Greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello, " + "Master!"
}

func Push(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Microsecond)
	}
}
